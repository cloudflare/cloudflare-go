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

// DLPDataTagCategoryDataTagService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDLPDataTagCategoryDataTagService] method instead.
type DLPDataTagCategoryDataTagService struct {
	Options []option.RequestOption
}

// NewDLPDataTagCategoryDataTagService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDLPDataTagCategoryDataTagService(opts ...option.RequestOption) (r *DLPDataTagCategoryDataTagService) {
	r = &DLPDataTagCategoryDataTagService{}
	r.Options = opts
	return
}

// Creates a new data tag.
func (r *DLPDataTagCategoryDataTagService) New(ctx context.Context, categoryID string, params DLPDataTagCategoryDataTagNewParams, opts ...option.RequestOption) (res *DLPDataTagCategoryDataTagNewResponse, err error) {
	var env DLPDataTagCategoryDataTagNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if categoryID == "" {
		err = errors.New("missing required category_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dlp/data_tag_categories/%s/data_tags", params.AccountID, categoryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Update the attributes of a single data tag.
func (r *DLPDataTagCategoryDataTagService) Update(ctx context.Context, categoryID string, tagID string, params DLPDataTagCategoryDataTagUpdateParams, opts ...option.RequestOption) (res *DLPDataTagCategoryDataTagUpdateResponse, err error) {
	var env DLPDataTagCategoryDataTagUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if categoryID == "" {
		err = errors.New("missing required category_id parameter")
		return nil, err
	}
	if tagID == "" {
		err = errors.New("missing required tag_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dlp/data_tag_categories/%s/data_tags/%s", params.AccountID, categoryID, tagID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieve all data tags in a data tag category
func (r *DLPDataTagCategoryDataTagService) List(ctx context.Context, categoryID string, query DLPDataTagCategoryDataTagListParams, opts ...option.RequestOption) (res *pagination.SinglePage[DLPDataTagCategoryDataTagListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if categoryID == "" {
		err = errors.New("missing required category_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dlp/data_tag_categories/%s/data_tags", query.AccountID, categoryID)
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

// Retrieve all data tags in a data tag category
func (r *DLPDataTagCategoryDataTagService) ListAutoPaging(ctx context.Context, categoryID string, query DLPDataTagCategoryDataTagListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[DLPDataTagCategoryDataTagListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, categoryID, query, opts...))
}

// Delete a single data tag.
func (r *DLPDataTagCategoryDataTagService) Delete(ctx context.Context, categoryID string, tagID string, body DLPDataTagCategoryDataTagDeleteParams, opts ...option.RequestOption) (res *DLPDataTagCategoryDataTagDeleteResponse, err error) {
	var env DLPDataTagCategoryDataTagDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if categoryID == "" {
		err = errors.New("missing required category_id parameter")
		return nil, err
	}
	if tagID == "" {
		err = errors.New("missing required tag_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dlp/data_tag_categories/%s/data_tags/%s", body.AccountID, categoryID, tagID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieve a specific data tag.
func (r *DLPDataTagCategoryDataTagService) Get(ctx context.Context, categoryID string, tagID string, query DLPDataTagCategoryDataTagGetParams, opts ...option.RequestOption) (res *DLPDataTagCategoryDataTagGetResponse, err error) {
	var env DLPDataTagCategoryDataTagGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if categoryID == "" {
		err = errors.New("missing required category_id parameter")
		return nil, err
	}
	if tagID == "" {
		err = errors.New("missing required tag_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dlp/data_tag_categories/%s/data_tags/%s", query.AccountID, categoryID, tagID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type DLPDataTagCategoryDataTagNewResponse struct {
	ID          string                                   `json:"id" api:"required" format:"uuid"`
	CreatedAt   time.Time                                `json:"created_at" api:"required" format:"date-time"`
	Name        string                                   `json:"name" api:"required"`
	UpdatedAt   time.Time                                `json:"updated_at" api:"required" format:"date-time"`
	Description string                                   `json:"description" api:"nullable"`
	JSON        dlpDataTagCategoryDataTagNewResponseJSON `json:"-"`
}

// dlpDataTagCategoryDataTagNewResponseJSON contains the JSON metadata for the
// struct [DLPDataTagCategoryDataTagNewResponse]
type dlpDataTagCategoryDataTagNewResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDataTagCategoryDataTagNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataTagCategoryDataTagNewResponseJSON) RawJSON() string {
	return r.raw
}

type DLPDataTagCategoryDataTagUpdateResponse struct {
	ID          string                                      `json:"id" api:"required" format:"uuid"`
	CreatedAt   time.Time                                   `json:"created_at" api:"required" format:"date-time"`
	Name        string                                      `json:"name" api:"required"`
	UpdatedAt   time.Time                                   `json:"updated_at" api:"required" format:"date-time"`
	Description string                                      `json:"description" api:"nullable"`
	JSON        dlpDataTagCategoryDataTagUpdateResponseJSON `json:"-"`
}

// dlpDataTagCategoryDataTagUpdateResponseJSON contains the JSON metadata for the
// struct [DLPDataTagCategoryDataTagUpdateResponse]
type dlpDataTagCategoryDataTagUpdateResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDataTagCategoryDataTagUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataTagCategoryDataTagUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type DLPDataTagCategoryDataTagListResponse struct {
	ID          string                                    `json:"id" api:"required" format:"uuid"`
	CreatedAt   time.Time                                 `json:"created_at" api:"required" format:"date-time"`
	Name        string                                    `json:"name" api:"required"`
	UpdatedAt   time.Time                                 `json:"updated_at" api:"required" format:"date-time"`
	Description string                                    `json:"description" api:"nullable"`
	JSON        dlpDataTagCategoryDataTagListResponseJSON `json:"-"`
}

// dlpDataTagCategoryDataTagListResponseJSON contains the JSON metadata for the
// struct [DLPDataTagCategoryDataTagListResponse]
type dlpDataTagCategoryDataTagListResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDataTagCategoryDataTagListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataTagCategoryDataTagListResponseJSON) RawJSON() string {
	return r.raw
}

type DLPDataTagCategoryDataTagDeleteResponse = interface{}

type DLPDataTagCategoryDataTagGetResponse struct {
	ID          string                                   `json:"id" api:"required" format:"uuid"`
	CreatedAt   time.Time                                `json:"created_at" api:"required" format:"date-time"`
	Name        string                                   `json:"name" api:"required"`
	UpdatedAt   time.Time                                `json:"updated_at" api:"required" format:"date-time"`
	Description string                                   `json:"description" api:"nullable"`
	JSON        dlpDataTagCategoryDataTagGetResponseJSON `json:"-"`
}

// dlpDataTagCategoryDataTagGetResponseJSON contains the JSON metadata for the
// struct [DLPDataTagCategoryDataTagGetResponse]
type dlpDataTagCategoryDataTagGetResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDataTagCategoryDataTagGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataTagCategoryDataTagGetResponseJSON) RawJSON() string {
	return r.raw
}

type DLPDataTagCategoryDataTagNewParams struct {
	AccountID   param.Field[string] `path:"account_id" api:"required"`
	Name        param.Field[string] `json:"name" api:"required"`
	Description param.Field[string] `json:"description"`
}

func (r DLPDataTagCategoryDataTagNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPDataTagCategoryDataTagNewResponseEnvelope struct {
	Errors   []DLPDataTagCategoryDataTagNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DLPDataTagCategoryDataTagNewResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success DLPDataTagCategoryDataTagNewResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  DLPDataTagCategoryDataTagNewResponse                `json:"result"`
	JSON    dlpDataTagCategoryDataTagNewResponseEnvelopeJSON    `json:"-"`
}

// dlpDataTagCategoryDataTagNewResponseEnvelopeJSON contains the JSON metadata for
// the struct [DLPDataTagCategoryDataTagNewResponseEnvelope]
type dlpDataTagCategoryDataTagNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDataTagCategoryDataTagNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataTagCategoryDataTagNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPDataTagCategoryDataTagNewResponseEnvelopeErrors struct {
	Code             int64                                                    `json:"code" api:"required"`
	Message          string                                                   `json:"message" api:"required"`
	DocumentationURL string                                                   `json:"documentation_url"`
	Source           DLPDataTagCategoryDataTagNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpDataTagCategoryDataTagNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpDataTagCategoryDataTagNewResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [DLPDataTagCategoryDataTagNewResponseEnvelopeErrors]
type dlpDataTagCategoryDataTagNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPDataTagCategoryDataTagNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataTagCategoryDataTagNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPDataTagCategoryDataTagNewResponseEnvelopeErrorsSource struct {
	Pointer string                                                       `json:"pointer"`
	JSON    dlpDataTagCategoryDataTagNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpDataTagCategoryDataTagNewResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [DLPDataTagCategoryDataTagNewResponseEnvelopeErrorsSource]
type dlpDataTagCategoryDataTagNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDataTagCategoryDataTagNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataTagCategoryDataTagNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPDataTagCategoryDataTagNewResponseEnvelopeMessages struct {
	Code             int64                                                      `json:"code" api:"required"`
	Message          string                                                     `json:"message" api:"required"`
	DocumentationURL string                                                     `json:"documentation_url"`
	Source           DLPDataTagCategoryDataTagNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpDataTagCategoryDataTagNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpDataTagCategoryDataTagNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [DLPDataTagCategoryDataTagNewResponseEnvelopeMessages]
type dlpDataTagCategoryDataTagNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPDataTagCategoryDataTagNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataTagCategoryDataTagNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPDataTagCategoryDataTagNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                         `json:"pointer"`
	JSON    dlpDataTagCategoryDataTagNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpDataTagCategoryDataTagNewResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [DLPDataTagCategoryDataTagNewResponseEnvelopeMessagesSource]
type dlpDataTagCategoryDataTagNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDataTagCategoryDataTagNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataTagCategoryDataTagNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// DLPDataTagCategoryDataTagNewResponseEnvelopeSuccess indicates whether the API call was successful.
type DLPDataTagCategoryDataTagNewResponseEnvelopeSuccess bool

const (
	DLPDataTagCategoryDataTagNewResponseEnvelopeSuccessTrue DLPDataTagCategoryDataTagNewResponseEnvelopeSuccess = true
)

func (r DLPDataTagCategoryDataTagNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPDataTagCategoryDataTagNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPDataTagCategoryDataTagUpdateParams struct {
	AccountID   param.Field[string] `path:"account_id" api:"required"`
	Description param.Field[string] `json:"description"`
	Name        param.Field[string] `json:"name"`
}

func (r DLPDataTagCategoryDataTagUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPDataTagCategoryDataTagUpdateResponseEnvelope struct {
	Errors   []DLPDataTagCategoryDataTagUpdateResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DLPDataTagCategoryDataTagUpdateResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success DLPDataTagCategoryDataTagUpdateResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  DLPDataTagCategoryDataTagUpdateResponse                `json:"result"`
	JSON    dlpDataTagCategoryDataTagUpdateResponseEnvelopeJSON    `json:"-"`
}

// dlpDataTagCategoryDataTagUpdateResponseEnvelopeJSON contains the JSON metadata
// for the struct [DLPDataTagCategoryDataTagUpdateResponseEnvelope]
type dlpDataTagCategoryDataTagUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDataTagCategoryDataTagUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataTagCategoryDataTagUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPDataTagCategoryDataTagUpdateResponseEnvelopeErrors struct {
	Code             int64                                                       `json:"code" api:"required"`
	Message          string                                                      `json:"message" api:"required"`
	DocumentationURL string                                                      `json:"documentation_url"`
	Source           DLPDataTagCategoryDataTagUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpDataTagCategoryDataTagUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpDataTagCategoryDataTagUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [DLPDataTagCategoryDataTagUpdateResponseEnvelopeErrors]
type dlpDataTagCategoryDataTagUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPDataTagCategoryDataTagUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataTagCategoryDataTagUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPDataTagCategoryDataTagUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                                          `json:"pointer"`
	JSON    dlpDataTagCategoryDataTagUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpDataTagCategoryDataTagUpdateResponseEnvelopeErrorsSourceJSON contains the
// JSON metadata for the struct
// [DLPDataTagCategoryDataTagUpdateResponseEnvelopeErrorsSource]
type dlpDataTagCategoryDataTagUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDataTagCategoryDataTagUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataTagCategoryDataTagUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPDataTagCategoryDataTagUpdateResponseEnvelopeMessages struct {
	Code             int64                                                         `json:"code" api:"required"`
	Message          string                                                        `json:"message" api:"required"`
	DocumentationURL string                                                        `json:"documentation_url"`
	Source           DLPDataTagCategoryDataTagUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpDataTagCategoryDataTagUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpDataTagCategoryDataTagUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [DLPDataTagCategoryDataTagUpdateResponseEnvelopeMessages]
type dlpDataTagCategoryDataTagUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPDataTagCategoryDataTagUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataTagCategoryDataTagUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPDataTagCategoryDataTagUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                            `json:"pointer"`
	JSON    dlpDataTagCategoryDataTagUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpDataTagCategoryDataTagUpdateResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [DLPDataTagCategoryDataTagUpdateResponseEnvelopeMessagesSource]
type dlpDataTagCategoryDataTagUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDataTagCategoryDataTagUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataTagCategoryDataTagUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// DLPDataTagCategoryDataTagUpdateResponseEnvelopeSuccess indicates whether the API call was successful.
type DLPDataTagCategoryDataTagUpdateResponseEnvelopeSuccess bool

const (
	DLPDataTagCategoryDataTagUpdateResponseEnvelopeSuccessTrue DLPDataTagCategoryDataTagUpdateResponseEnvelopeSuccess = true
)

func (r DLPDataTagCategoryDataTagUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPDataTagCategoryDataTagUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPDataTagCategoryDataTagListParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type DLPDataTagCategoryDataTagDeleteParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type DLPDataTagCategoryDataTagDeleteResponseEnvelope struct {
	Errors   []DLPDataTagCategoryDataTagDeleteResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DLPDataTagCategoryDataTagDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success DLPDataTagCategoryDataTagDeleteResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  DLPDataTagCategoryDataTagDeleteResponse                `json:"result" api:"nullable"`
	JSON    dlpDataTagCategoryDataTagDeleteResponseEnvelopeJSON    `json:"-"`
}

// dlpDataTagCategoryDataTagDeleteResponseEnvelopeJSON contains the JSON metadata
// for the struct [DLPDataTagCategoryDataTagDeleteResponseEnvelope]
type dlpDataTagCategoryDataTagDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDataTagCategoryDataTagDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataTagCategoryDataTagDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPDataTagCategoryDataTagDeleteResponseEnvelopeErrors struct {
	Code             int64                                                       `json:"code" api:"required"`
	Message          string                                                      `json:"message" api:"required"`
	DocumentationURL string                                                      `json:"documentation_url"`
	Source           DLPDataTagCategoryDataTagDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpDataTagCategoryDataTagDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpDataTagCategoryDataTagDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [DLPDataTagCategoryDataTagDeleteResponseEnvelopeErrors]
type dlpDataTagCategoryDataTagDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPDataTagCategoryDataTagDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataTagCategoryDataTagDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPDataTagCategoryDataTagDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                          `json:"pointer"`
	JSON    dlpDataTagCategoryDataTagDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpDataTagCategoryDataTagDeleteResponseEnvelopeErrorsSourceJSON contains the
// JSON metadata for the struct
// [DLPDataTagCategoryDataTagDeleteResponseEnvelopeErrorsSource]
type dlpDataTagCategoryDataTagDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDataTagCategoryDataTagDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataTagCategoryDataTagDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPDataTagCategoryDataTagDeleteResponseEnvelopeMessages struct {
	Code             int64                                                         `json:"code" api:"required"`
	Message          string                                                        `json:"message" api:"required"`
	DocumentationURL string                                                        `json:"documentation_url"`
	Source           DLPDataTagCategoryDataTagDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpDataTagCategoryDataTagDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpDataTagCategoryDataTagDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [DLPDataTagCategoryDataTagDeleteResponseEnvelopeMessages]
type dlpDataTagCategoryDataTagDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPDataTagCategoryDataTagDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataTagCategoryDataTagDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPDataTagCategoryDataTagDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                            `json:"pointer"`
	JSON    dlpDataTagCategoryDataTagDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpDataTagCategoryDataTagDeleteResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [DLPDataTagCategoryDataTagDeleteResponseEnvelopeMessagesSource]
type dlpDataTagCategoryDataTagDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDataTagCategoryDataTagDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataTagCategoryDataTagDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// DLPDataTagCategoryDataTagDeleteResponseEnvelopeSuccess indicates whether the API call was successful.
type DLPDataTagCategoryDataTagDeleteResponseEnvelopeSuccess bool

const (
	DLPDataTagCategoryDataTagDeleteResponseEnvelopeSuccessTrue DLPDataTagCategoryDataTagDeleteResponseEnvelopeSuccess = true
)

func (r DLPDataTagCategoryDataTagDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPDataTagCategoryDataTagDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPDataTagCategoryDataTagGetParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type DLPDataTagCategoryDataTagGetResponseEnvelope struct {
	Errors   []DLPDataTagCategoryDataTagGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DLPDataTagCategoryDataTagGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success DLPDataTagCategoryDataTagGetResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  DLPDataTagCategoryDataTagGetResponse                `json:"result"`
	JSON    dlpDataTagCategoryDataTagGetResponseEnvelopeJSON    `json:"-"`
}

// dlpDataTagCategoryDataTagGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [DLPDataTagCategoryDataTagGetResponseEnvelope]
type dlpDataTagCategoryDataTagGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDataTagCategoryDataTagGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataTagCategoryDataTagGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPDataTagCategoryDataTagGetResponseEnvelopeErrors struct {
	Code             int64                                                    `json:"code" api:"required"`
	Message          string                                                   `json:"message" api:"required"`
	DocumentationURL string                                                   `json:"documentation_url"`
	Source           DLPDataTagCategoryDataTagGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpDataTagCategoryDataTagGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpDataTagCategoryDataTagGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [DLPDataTagCategoryDataTagGetResponseEnvelopeErrors]
type dlpDataTagCategoryDataTagGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPDataTagCategoryDataTagGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataTagCategoryDataTagGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPDataTagCategoryDataTagGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                       `json:"pointer"`
	JSON    dlpDataTagCategoryDataTagGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpDataTagCategoryDataTagGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [DLPDataTagCategoryDataTagGetResponseEnvelopeErrorsSource]
type dlpDataTagCategoryDataTagGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDataTagCategoryDataTagGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataTagCategoryDataTagGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPDataTagCategoryDataTagGetResponseEnvelopeMessages struct {
	Code             int64                                                      `json:"code" api:"required"`
	Message          string                                                     `json:"message" api:"required"`
	DocumentationURL string                                                     `json:"documentation_url"`
	Source           DLPDataTagCategoryDataTagGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpDataTagCategoryDataTagGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpDataTagCategoryDataTagGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [DLPDataTagCategoryDataTagGetResponseEnvelopeMessages]
type dlpDataTagCategoryDataTagGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPDataTagCategoryDataTagGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataTagCategoryDataTagGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPDataTagCategoryDataTagGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                         `json:"pointer"`
	JSON    dlpDataTagCategoryDataTagGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpDataTagCategoryDataTagGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [DLPDataTagCategoryDataTagGetResponseEnvelopeMessagesSource]
type dlpDataTagCategoryDataTagGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDataTagCategoryDataTagGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataTagCategoryDataTagGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// DLPDataTagCategoryDataTagGetResponseEnvelopeSuccess indicates whether the API call was successful.
type DLPDataTagCategoryDataTagGetResponseEnvelopeSuccess bool

const (
	DLPDataTagCategoryDataTagGetResponseEnvelopeSuccessTrue DLPDataTagCategoryDataTagGetResponseEnvelopeSuccess = true
)

func (r DLPDataTagCategoryDataTagGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPDataTagCategoryDataTagGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
