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

// DLPCustomPromptTopicService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDLPCustomPromptTopicService] method instead.
type DLPCustomPromptTopicService struct {
	Options []option.RequestOption
}

// NewDLPCustomPromptTopicService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDLPCustomPromptTopicService(opts ...option.RequestOption) (r *DLPCustomPromptTopicService) {
	r = &DLPCustomPromptTopicService{}
	r.Options = opts
	return
}

// Creates a DLP custom prompt topic entry.
func (r *DLPCustomPromptTopicService) New(ctx context.Context, params DLPCustomPromptTopicNewParams, opts ...option.RequestOption) (res *CustomPromptTopic, err error) {
	var env DLPCustomPromptTopicNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dlp/custom_prompt_topics", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Updates a DLP custom prompt topic entry.
func (r *DLPCustomPromptTopicService) Update(ctx context.Context, entryID string, params DLPCustomPromptTopicUpdateParams, opts ...option.RequestOption) (res *CustomPromptTopic, err error) {
	var env DLPCustomPromptTopicUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if entryID == "" {
		err = errors.New("missing required entry_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dlp/custom_prompt_topics/%s", params.AccountID, entryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Lists all DLP custom prompt topic entries in an account.
func (r *DLPCustomPromptTopicService) List(ctx context.Context, query DLPCustomPromptTopicListParams, opts ...option.RequestOption) (res *pagination.SinglePage[CustomPromptTopic], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dlp/custom_prompt_topics", query.AccountID)
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

// Lists all DLP custom prompt topic entries in an account.
func (r *DLPCustomPromptTopicService) ListAutoPaging(ctx context.Context, query DLPCustomPromptTopicListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[CustomPromptTopic] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Deletes a DLP custom prompt topic entry.
func (r *DLPCustomPromptTopicService) Delete(ctx context.Context, entryID string, body DLPCustomPromptTopicDeleteParams, opts ...option.RequestOption) (res *DLPCustomPromptTopicDeleteResponse, err error) {
	var env DLPCustomPromptTopicDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if entryID == "" {
		err = errors.New("missing required entry_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dlp/custom_prompt_topics/%s", body.AccountID, entryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Fetches a DLP custom prompt topic entry by ID.
func (r *DLPCustomPromptTopicService) Get(ctx context.Context, entryID string, query DLPCustomPromptTopicGetParams, opts ...option.RequestOption) (res *CustomPromptTopic, err error) {
	var env DLPCustomPromptTopicGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if entryID == "" {
		err = errors.New("missing required entry_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dlp/custom_prompt_topics/%s", query.AccountID, entryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type CustomPromptTopic struct {
	ID        string    `json:"id" api:"required" format:"uuid"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Deprecated: deprecated
	Enabled     bool      `json:"enabled" api:"required"`
	Name        string    `json:"name" api:"required"`
	Topic       string    `json:"topic" api:"required"`
	UpdatedAt   time.Time `json:"updated_at" api:"required" format:"date-time"`
	Description string    `json:"description" api:"nullable"`
	// Deprecated: deprecated
	ProfileID string                `json:"profile_id" api:"nullable" format:"uuid"`
	JSON      customPromptTopicJSON `json:"-"`
}

// customPromptTopicJSON contains the JSON metadata for the struct
// [CustomPromptTopic]
type customPromptTopicJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Topic       apijson.Field
	UpdatedAt   apijson.Field
	Description apijson.Field
	ProfileID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomPromptTopic) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customPromptTopicJSON) RawJSON() string {
	return r.raw
}

type DLPCustomPromptTopicDeleteResponse = interface{}

type DLPCustomPromptTopicNewParams struct {
	AccountID   param.Field[string] `path:"account_id" api:"required"`
	Enabled     param.Field[bool]   `json:"enabled" api:"required"`
	Name        param.Field[string] `json:"name" api:"required"`
	Topic       param.Field[string] `json:"topic" api:"required"`
	Description param.Field[string] `json:"description"`
	ProfileID   param.Field[string] `json:"profile_id" format:"uuid"`
}

func (r DLPCustomPromptTopicNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPCustomPromptTopicNewResponseEnvelope struct {
	Errors   []DLPCustomPromptTopicNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DLPCustomPromptTopicNewResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success DLPCustomPromptTopicNewResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  CustomPromptTopic                              `json:"result"`
	JSON    dlpCustomPromptTopicNewResponseEnvelopeJSON    `json:"-"`
}

// dlpCustomPromptTopicNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPCustomPromptTopicNewResponseEnvelope]
type dlpCustomPromptTopicNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPCustomPromptTopicNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpCustomPromptTopicNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPCustomPromptTopicNewResponseEnvelopeErrors struct {
	Code             int64                                               `json:"code" api:"required"`
	Message          string                                              `json:"message" api:"required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           DLPCustomPromptTopicNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpCustomPromptTopicNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpCustomPromptTopicNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DLPCustomPromptTopicNewResponseEnvelopeErrors]
type dlpCustomPromptTopicNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPCustomPromptTopicNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpCustomPromptTopicNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPCustomPromptTopicNewResponseEnvelopeErrorsSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    dlpCustomPromptTopicNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpCustomPromptTopicNewResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [DLPCustomPromptTopicNewResponseEnvelopeErrorsSource]
type dlpCustomPromptTopicNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPCustomPromptTopicNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpCustomPromptTopicNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPCustomPromptTopicNewResponseEnvelopeMessages struct {
	Code             int64                                                 `json:"code" api:"required"`
	Message          string                                                `json:"message" api:"required"`
	DocumentationURL string                                                `json:"documentation_url"`
	Source           DLPCustomPromptTopicNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpCustomPromptTopicNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpCustomPromptTopicNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DLPCustomPromptTopicNewResponseEnvelopeMessages]
type dlpCustomPromptTopicNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPCustomPromptTopicNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpCustomPromptTopicNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPCustomPromptTopicNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                    `json:"pointer"`
	JSON    dlpCustomPromptTopicNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpCustomPromptTopicNewResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [DLPCustomPromptTopicNewResponseEnvelopeMessagesSource]
type dlpCustomPromptTopicNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPCustomPromptTopicNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpCustomPromptTopicNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPCustomPromptTopicNewResponseEnvelopeSuccess bool

const (
	DLPCustomPromptTopicNewResponseEnvelopeSuccessTrue DLPCustomPromptTopicNewResponseEnvelopeSuccess = true
)

func (r DLPCustomPromptTopicNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPCustomPromptTopicNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPCustomPromptTopicUpdateParams struct {
	AccountID   param.Field[string] `path:"account_id" api:"required"`
	Enabled     param.Field[bool]   `json:"enabled" api:"required"`
	Name        param.Field[string] `json:"name" api:"required"`
	Topic       param.Field[string] `json:"topic" api:"required"`
	Description param.Field[string] `json:"description"`
}

func (r DLPCustomPromptTopicUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPCustomPromptTopicUpdateResponseEnvelope struct {
	Errors   []DLPCustomPromptTopicUpdateResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DLPCustomPromptTopicUpdateResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success DLPCustomPromptTopicUpdateResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  CustomPromptTopic                                 `json:"result"`
	JSON    dlpCustomPromptTopicUpdateResponseEnvelopeJSON    `json:"-"`
}

// dlpCustomPromptTopicUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [DLPCustomPromptTopicUpdateResponseEnvelope]
type dlpCustomPromptTopicUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPCustomPromptTopicUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpCustomPromptTopicUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPCustomPromptTopicUpdateResponseEnvelopeErrors struct {
	Code             int64                                                  `json:"code" api:"required"`
	Message          string                                                 `json:"message" api:"required"`
	DocumentationURL string                                                 `json:"documentation_url"`
	Source           DLPCustomPromptTopicUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpCustomPromptTopicUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpCustomPromptTopicUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [DLPCustomPromptTopicUpdateResponseEnvelopeErrors]
type dlpCustomPromptTopicUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPCustomPromptTopicUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpCustomPromptTopicUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPCustomPromptTopicUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                                     `json:"pointer"`
	JSON    dlpCustomPromptTopicUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpCustomPromptTopicUpdateResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [DLPCustomPromptTopicUpdateResponseEnvelopeErrorsSource]
type dlpCustomPromptTopicUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPCustomPromptTopicUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpCustomPromptTopicUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPCustomPromptTopicUpdateResponseEnvelopeMessages struct {
	Code             int64                                                    `json:"code" api:"required"`
	Message          string                                                   `json:"message" api:"required"`
	DocumentationURL string                                                   `json:"documentation_url"`
	Source           DLPCustomPromptTopicUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpCustomPromptTopicUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpCustomPromptTopicUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [DLPCustomPromptTopicUpdateResponseEnvelopeMessages]
type dlpCustomPromptTopicUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPCustomPromptTopicUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpCustomPromptTopicUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPCustomPromptTopicUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                       `json:"pointer"`
	JSON    dlpCustomPromptTopicUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpCustomPromptTopicUpdateResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [DLPCustomPromptTopicUpdateResponseEnvelopeMessagesSource]
type dlpCustomPromptTopicUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPCustomPromptTopicUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpCustomPromptTopicUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPCustomPromptTopicUpdateResponseEnvelopeSuccess bool

const (
	DLPCustomPromptTopicUpdateResponseEnvelopeSuccessTrue DLPCustomPromptTopicUpdateResponseEnvelopeSuccess = true
)

func (r DLPCustomPromptTopicUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPCustomPromptTopicUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPCustomPromptTopicListParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type DLPCustomPromptTopicDeleteParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type DLPCustomPromptTopicDeleteResponseEnvelope struct {
	Errors   []DLPCustomPromptTopicDeleteResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DLPCustomPromptTopicDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success DLPCustomPromptTopicDeleteResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  DLPCustomPromptTopicDeleteResponse                `json:"result" api:"nullable"`
	JSON    dlpCustomPromptTopicDeleteResponseEnvelopeJSON    `json:"-"`
}

// dlpCustomPromptTopicDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [DLPCustomPromptTopicDeleteResponseEnvelope]
type dlpCustomPromptTopicDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPCustomPromptTopicDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpCustomPromptTopicDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPCustomPromptTopicDeleteResponseEnvelopeErrors struct {
	Code             int64                                                  `json:"code" api:"required"`
	Message          string                                                 `json:"message" api:"required"`
	DocumentationURL string                                                 `json:"documentation_url"`
	Source           DLPCustomPromptTopicDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpCustomPromptTopicDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpCustomPromptTopicDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [DLPCustomPromptTopicDeleteResponseEnvelopeErrors]
type dlpCustomPromptTopicDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPCustomPromptTopicDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpCustomPromptTopicDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPCustomPromptTopicDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                     `json:"pointer"`
	JSON    dlpCustomPromptTopicDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpCustomPromptTopicDeleteResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [DLPCustomPromptTopicDeleteResponseEnvelopeErrorsSource]
type dlpCustomPromptTopicDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPCustomPromptTopicDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpCustomPromptTopicDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPCustomPromptTopicDeleteResponseEnvelopeMessages struct {
	Code             int64                                                    `json:"code" api:"required"`
	Message          string                                                   `json:"message" api:"required"`
	DocumentationURL string                                                   `json:"documentation_url"`
	Source           DLPCustomPromptTopicDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpCustomPromptTopicDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpCustomPromptTopicDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [DLPCustomPromptTopicDeleteResponseEnvelopeMessages]
type dlpCustomPromptTopicDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPCustomPromptTopicDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpCustomPromptTopicDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPCustomPromptTopicDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                       `json:"pointer"`
	JSON    dlpCustomPromptTopicDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpCustomPromptTopicDeleteResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [DLPCustomPromptTopicDeleteResponseEnvelopeMessagesSource]
type dlpCustomPromptTopicDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPCustomPromptTopicDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpCustomPromptTopicDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPCustomPromptTopicDeleteResponseEnvelopeSuccess bool

const (
	DLPCustomPromptTopicDeleteResponseEnvelopeSuccessTrue DLPCustomPromptTopicDeleteResponseEnvelopeSuccess = true
)

func (r DLPCustomPromptTopicDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPCustomPromptTopicDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPCustomPromptTopicGetParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type DLPCustomPromptTopicGetResponseEnvelope struct {
	Errors   []DLPCustomPromptTopicGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DLPCustomPromptTopicGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success DLPCustomPromptTopicGetResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  CustomPromptTopic                              `json:"result"`
	JSON    dlpCustomPromptTopicGetResponseEnvelopeJSON    `json:"-"`
}

// dlpCustomPromptTopicGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPCustomPromptTopicGetResponseEnvelope]
type dlpCustomPromptTopicGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPCustomPromptTopicGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpCustomPromptTopicGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPCustomPromptTopicGetResponseEnvelopeErrors struct {
	Code             int64                                               `json:"code" api:"required"`
	Message          string                                              `json:"message" api:"required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           DLPCustomPromptTopicGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpCustomPromptTopicGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpCustomPromptTopicGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DLPCustomPromptTopicGetResponseEnvelopeErrors]
type dlpCustomPromptTopicGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPCustomPromptTopicGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpCustomPromptTopicGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPCustomPromptTopicGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    dlpCustomPromptTopicGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpCustomPromptTopicGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [DLPCustomPromptTopicGetResponseEnvelopeErrorsSource]
type dlpCustomPromptTopicGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPCustomPromptTopicGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpCustomPromptTopicGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPCustomPromptTopicGetResponseEnvelopeMessages struct {
	Code             int64                                                 `json:"code" api:"required"`
	Message          string                                                `json:"message" api:"required"`
	DocumentationURL string                                                `json:"documentation_url"`
	Source           DLPCustomPromptTopicGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpCustomPromptTopicGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpCustomPromptTopicGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DLPCustomPromptTopicGetResponseEnvelopeMessages]
type dlpCustomPromptTopicGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPCustomPromptTopicGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpCustomPromptTopicGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPCustomPromptTopicGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                    `json:"pointer"`
	JSON    dlpCustomPromptTopicGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpCustomPromptTopicGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [DLPCustomPromptTopicGetResponseEnvelopeMessagesSource]
type dlpCustomPromptTopicGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPCustomPromptTopicGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpCustomPromptTopicGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPCustomPromptTopicGetResponseEnvelopeSuccess bool

const (
	DLPCustomPromptTopicGetResponseEnvelopeSuccessTrue DLPCustomPromptTopicGetResponseEnvelopeSuccess = true
)

func (r DLPCustomPromptTopicGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPCustomPromptTopicGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
