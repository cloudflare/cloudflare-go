// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
)

// SettingContentPolicyService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingContentPolicyService] method instead.
type SettingContentPolicyService struct {
	Options []option.RequestOption
}

// NewSettingContentPolicyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingContentPolicyService(opts ...option.RequestOption) (r *SettingContentPolicyService) {
	r = &SettingContentPolicyService{}
	r.Options = opts
	return
}

// Creates a new content policy. Emails whose subject or body matches the pattern
// will be subject to the configured action.
func (r *SettingContentPolicyService) New(ctx context.Context, params SettingContentPolicyNewParams, opts ...option.RequestOption) (res *SettingContentPolicyNewResponse, err error) {
	var env SettingContentPolicyNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/content_policies", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns a paginated list of content policies. These policies match against the
// subject or body of emails using a pattern. Supports filtering by name or enabled
// status, and searching across name and pattern fields.
func (r *SettingContentPolicyService) List(ctx context.Context, params SettingContentPolicyListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[SettingContentPolicyListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/content_policies", params.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// Returns a paginated list of content policies. These policies match against the
// subject or body of emails using a pattern. Supports filtering by name or enabled
// status, and searching across name and pattern fields.
func (r *SettingContentPolicyService) ListAutoPaging(ctx context.Context, params SettingContentPolicyListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[SettingContentPolicyListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Removes a content policy. After deletion, emails will no longer be evaluated
// against this pattern.
func (r *SettingContentPolicyService) Delete(ctx context.Context, policyID string, body SettingContentPolicyDeleteParams, opts ...option.RequestOption) (res *SettingContentPolicyDeleteResponse, err error) {
	var env SettingContentPolicyDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/content_policies/%s", body.AccountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Executes multiple operations atomically. All four operation arrays (deletes,
// patches, puts, posts) are required and executed in order. Send empty arrays for
// unused operations.
func (r *SettingContentPolicyService) Batch(ctx context.Context, params SettingContentPolicyBatchParams, opts ...option.RequestOption) (res *SettingContentPolicyBatchResponse, err error) {
	var env SettingContentPolicyBatchResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/content_policies/batch", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Updates an existing content policy. Only provided fields will be modified.
func (r *SettingContentPolicyService) Edit(ctx context.Context, policyID string, params SettingContentPolicyEditParams, opts ...option.RequestOption) (res *SettingContentPolicyEditResponse, err error) {
	var env SettingContentPolicyEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/content_policies/%s", params.AccountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieves details for a specific content policy including its pattern, targets,
// and metadata.
func (r *SettingContentPolicyService) Get(ctx context.Context, policyID string, query SettingContentPolicyGetParams, opts ...option.RequestOption) (res *SettingContentPolicyGetResponse, err error) {
	var env SettingContentPolicyGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/content_policies/%s", query.AccountID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// A content policy pattern that matches against the subject or body of an email.
type SettingContentPolicyNewResponse struct {
	// Content policy identifier.
	ID         string                                  `json:"id" format:"uuid"`
	CreatedAt  time.Time                               `json:"created_at" format:"date-time"`
	Enabled    bool                                    `json:"enabled"`
	ModifiedAt time.Time                               `json:"modified_at" format:"date-time"`
	Name       string                                  `json:"name"`
	Notes      string                                  `json:"notes" api:"nullable"`
	Pattern    string                                  `json:"pattern"`
	Targets    []SettingContentPolicyNewResponseTarget `json:"targets"`
	JSON       settingContentPolicyNewResponseJSON     `json:"-"`
}

// settingContentPolicyNewResponseJSON contains the JSON metadata for the struct
// [SettingContentPolicyNewResponse]
type settingContentPolicyNewResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	Notes       apijson.Field
	Pattern     apijson.Field
	Targets     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingContentPolicyNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingContentPolicyNewResponseJSON) RawJSON() string {
	return r.raw
}

// The part of the email to match the pattern against.
type SettingContentPolicyNewResponseTarget string

const (
	SettingContentPolicyNewResponseTargetSubject SettingContentPolicyNewResponseTarget = "SUBJECT"
	SettingContentPolicyNewResponseTargetBody    SettingContentPolicyNewResponseTarget = "BODY"
)

func (r SettingContentPolicyNewResponseTarget) IsKnown() bool {
	switch r {
	case SettingContentPolicyNewResponseTargetSubject, SettingContentPolicyNewResponseTargetBody:
		return true
	}
	return false
}

// A content policy pattern that matches against the subject or body of an email.
type SettingContentPolicyListResponse struct {
	// Content policy identifier.
	ID         string                                   `json:"id" format:"uuid"`
	CreatedAt  time.Time                                `json:"created_at" format:"date-time"`
	Enabled    bool                                     `json:"enabled"`
	ModifiedAt time.Time                                `json:"modified_at" format:"date-time"`
	Name       string                                   `json:"name"`
	Notes      string                                   `json:"notes" api:"nullable"`
	Pattern    string                                   `json:"pattern"`
	Targets    []SettingContentPolicyListResponseTarget `json:"targets"`
	JSON       settingContentPolicyListResponseJSON     `json:"-"`
}

// settingContentPolicyListResponseJSON contains the JSON metadata for the struct
// [SettingContentPolicyListResponse]
type settingContentPolicyListResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	Notes       apijson.Field
	Pattern     apijson.Field
	Targets     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingContentPolicyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingContentPolicyListResponseJSON) RawJSON() string {
	return r.raw
}

// The part of the email to match the pattern against.
type SettingContentPolicyListResponseTarget string

const (
	SettingContentPolicyListResponseTargetSubject SettingContentPolicyListResponseTarget = "SUBJECT"
	SettingContentPolicyListResponseTargetBody    SettingContentPolicyListResponseTarget = "BODY"
)

func (r SettingContentPolicyListResponseTarget) IsKnown() bool {
	switch r {
	case SettingContentPolicyListResponseTargetSubject, SettingContentPolicyListResponseTargetBody:
		return true
	}
	return false
}

type SettingContentPolicyDeleteResponse struct {
	// Content policy identifier.
	ID   string                                 `json:"id" api:"required" format:"uuid"`
	JSON settingContentPolicyDeleteResponseJSON `json:"-"`
}

// settingContentPolicyDeleteResponseJSON contains the JSON metadata for the struct
// [SettingContentPolicyDeleteResponse]
type settingContentPolicyDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingContentPolicyDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingContentPolicyDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type SettingContentPolicyBatchResponse struct {
	Deletes []SettingContentPolicyBatchResponseDelete `json:"deletes"`
	Patches []SettingContentPolicyBatchResponsePatch  `json:"patches"`
	Posts   []SettingContentPolicyBatchResponsePost   `json:"posts"`
	Puts    []SettingContentPolicyBatchResponsePut    `json:"puts"`
	JSON    settingContentPolicyBatchResponseJSON     `json:"-"`
}

// settingContentPolicyBatchResponseJSON contains the JSON metadata for the struct
// [SettingContentPolicyBatchResponse]
type settingContentPolicyBatchResponseJSON struct {
	Deletes     apijson.Field
	Patches     apijson.Field
	Posts       apijson.Field
	Puts        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingContentPolicyBatchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingContentPolicyBatchResponseJSON) RawJSON() string {
	return r.raw
}

type SettingContentPolicyBatchResponseDelete struct {
	// Content policy identifier.
	ID   string                                      `json:"id" api:"required" format:"uuid"`
	JSON settingContentPolicyBatchResponseDeleteJSON `json:"-"`
}

// settingContentPolicyBatchResponseDeleteJSON contains the JSON metadata for the
// struct [SettingContentPolicyBatchResponseDelete]
type settingContentPolicyBatchResponseDeleteJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingContentPolicyBatchResponseDelete) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingContentPolicyBatchResponseDeleteJSON) RawJSON() string {
	return r.raw
}

// A content policy pattern that matches against the subject or body of an email.
type SettingContentPolicyBatchResponsePatch struct {
	// Content policy identifier.
	ID         string                                           `json:"id" format:"uuid"`
	CreatedAt  time.Time                                        `json:"created_at" format:"date-time"`
	Enabled    bool                                             `json:"enabled"`
	ModifiedAt time.Time                                        `json:"modified_at" format:"date-time"`
	Name       string                                           `json:"name"`
	Notes      string                                           `json:"notes" api:"nullable"`
	Pattern    string                                           `json:"pattern"`
	Targets    []SettingContentPolicyBatchResponsePatchesTarget `json:"targets"`
	JSON       settingContentPolicyBatchResponsePatchJSON       `json:"-"`
}

// settingContentPolicyBatchResponsePatchJSON contains the JSON metadata for the
// struct [SettingContentPolicyBatchResponsePatch]
type settingContentPolicyBatchResponsePatchJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	Notes       apijson.Field
	Pattern     apijson.Field
	Targets     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingContentPolicyBatchResponsePatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingContentPolicyBatchResponsePatchJSON) RawJSON() string {
	return r.raw
}

// The part of the email to match the pattern against.
type SettingContentPolicyBatchResponsePatchesTarget string

const (
	SettingContentPolicyBatchResponsePatchesTargetSubject SettingContentPolicyBatchResponsePatchesTarget = "SUBJECT"
	SettingContentPolicyBatchResponsePatchesTargetBody    SettingContentPolicyBatchResponsePatchesTarget = "BODY"
)

func (r SettingContentPolicyBatchResponsePatchesTarget) IsKnown() bool {
	switch r {
	case SettingContentPolicyBatchResponsePatchesTargetSubject, SettingContentPolicyBatchResponsePatchesTargetBody:
		return true
	}
	return false
}

// A content policy pattern that matches against the subject or body of an email.
type SettingContentPolicyBatchResponsePost struct {
	// Content policy identifier.
	ID         string                                         `json:"id" format:"uuid"`
	CreatedAt  time.Time                                      `json:"created_at" format:"date-time"`
	Enabled    bool                                           `json:"enabled"`
	ModifiedAt time.Time                                      `json:"modified_at" format:"date-time"`
	Name       string                                         `json:"name"`
	Notes      string                                         `json:"notes" api:"nullable"`
	Pattern    string                                         `json:"pattern"`
	Targets    []SettingContentPolicyBatchResponsePostsTarget `json:"targets"`
	JSON       settingContentPolicyBatchResponsePostJSON      `json:"-"`
}

// settingContentPolicyBatchResponsePostJSON contains the JSON metadata for the
// struct [SettingContentPolicyBatchResponsePost]
type settingContentPolicyBatchResponsePostJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	Notes       apijson.Field
	Pattern     apijson.Field
	Targets     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingContentPolicyBatchResponsePost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingContentPolicyBatchResponsePostJSON) RawJSON() string {
	return r.raw
}

// The part of the email to match the pattern against.
type SettingContentPolicyBatchResponsePostsTarget string

const (
	SettingContentPolicyBatchResponsePostsTargetSubject SettingContentPolicyBatchResponsePostsTarget = "SUBJECT"
	SettingContentPolicyBatchResponsePostsTargetBody    SettingContentPolicyBatchResponsePostsTarget = "BODY"
)

func (r SettingContentPolicyBatchResponsePostsTarget) IsKnown() bool {
	switch r {
	case SettingContentPolicyBatchResponsePostsTargetSubject, SettingContentPolicyBatchResponsePostsTargetBody:
		return true
	}
	return false
}

// A content policy pattern that matches against the subject or body of an email.
type SettingContentPolicyBatchResponsePut struct {
	// Content policy identifier.
	ID         string                                        `json:"id" format:"uuid"`
	CreatedAt  time.Time                                     `json:"created_at" format:"date-time"`
	Enabled    bool                                          `json:"enabled"`
	ModifiedAt time.Time                                     `json:"modified_at" format:"date-time"`
	Name       string                                        `json:"name"`
	Notes      string                                        `json:"notes" api:"nullable"`
	Pattern    string                                        `json:"pattern"`
	Targets    []SettingContentPolicyBatchResponsePutsTarget `json:"targets"`
	JSON       settingContentPolicyBatchResponsePutJSON      `json:"-"`
}

// settingContentPolicyBatchResponsePutJSON contains the JSON metadata for the
// struct [SettingContentPolicyBatchResponsePut]
type settingContentPolicyBatchResponsePutJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	Notes       apijson.Field
	Pattern     apijson.Field
	Targets     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingContentPolicyBatchResponsePut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingContentPolicyBatchResponsePutJSON) RawJSON() string {
	return r.raw
}

// The part of the email to match the pattern against.
type SettingContentPolicyBatchResponsePutsTarget string

const (
	SettingContentPolicyBatchResponsePutsTargetSubject SettingContentPolicyBatchResponsePutsTarget = "SUBJECT"
	SettingContentPolicyBatchResponsePutsTargetBody    SettingContentPolicyBatchResponsePutsTarget = "BODY"
)

func (r SettingContentPolicyBatchResponsePutsTarget) IsKnown() bool {
	switch r {
	case SettingContentPolicyBatchResponsePutsTargetSubject, SettingContentPolicyBatchResponsePutsTargetBody:
		return true
	}
	return false
}

// A content policy pattern that matches against the subject or body of an email.
type SettingContentPolicyEditResponse struct {
	// Content policy identifier.
	ID         string                                   `json:"id" format:"uuid"`
	CreatedAt  time.Time                                `json:"created_at" format:"date-time"`
	Enabled    bool                                     `json:"enabled"`
	ModifiedAt time.Time                                `json:"modified_at" format:"date-time"`
	Name       string                                   `json:"name"`
	Notes      string                                   `json:"notes" api:"nullable"`
	Pattern    string                                   `json:"pattern"`
	Targets    []SettingContentPolicyEditResponseTarget `json:"targets"`
	JSON       settingContentPolicyEditResponseJSON     `json:"-"`
}

// settingContentPolicyEditResponseJSON contains the JSON metadata for the struct
// [SettingContentPolicyEditResponse]
type settingContentPolicyEditResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	Notes       apijson.Field
	Pattern     apijson.Field
	Targets     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingContentPolicyEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingContentPolicyEditResponseJSON) RawJSON() string {
	return r.raw
}

// The part of the email to match the pattern against.
type SettingContentPolicyEditResponseTarget string

const (
	SettingContentPolicyEditResponseTargetSubject SettingContentPolicyEditResponseTarget = "SUBJECT"
	SettingContentPolicyEditResponseTargetBody    SettingContentPolicyEditResponseTarget = "BODY"
)

func (r SettingContentPolicyEditResponseTarget) IsKnown() bool {
	switch r {
	case SettingContentPolicyEditResponseTargetSubject, SettingContentPolicyEditResponseTargetBody:
		return true
	}
	return false
}

// A content policy pattern that matches against the subject or body of an email.
type SettingContentPolicyGetResponse struct {
	// Content policy identifier.
	ID         string                                  `json:"id" format:"uuid"`
	CreatedAt  time.Time                               `json:"created_at" format:"date-time"`
	Enabled    bool                                    `json:"enabled"`
	ModifiedAt time.Time                               `json:"modified_at" format:"date-time"`
	Name       string                                  `json:"name"`
	Notes      string                                  `json:"notes" api:"nullable"`
	Pattern    string                                  `json:"pattern"`
	Targets    []SettingContentPolicyGetResponseTarget `json:"targets"`
	JSON       settingContentPolicyGetResponseJSON     `json:"-"`
}

// settingContentPolicyGetResponseJSON contains the JSON metadata for the struct
// [SettingContentPolicyGetResponse]
type settingContentPolicyGetResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	Notes       apijson.Field
	Pattern     apijson.Field
	Targets     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingContentPolicyGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingContentPolicyGetResponseJSON) RawJSON() string {
	return r.raw
}

// The part of the email to match the pattern against.
type SettingContentPolicyGetResponseTarget string

const (
	SettingContentPolicyGetResponseTargetSubject SettingContentPolicyGetResponseTarget = "SUBJECT"
	SettingContentPolicyGetResponseTargetBody    SettingContentPolicyGetResponseTarget = "BODY"
)

func (r SettingContentPolicyGetResponseTarget) IsKnown() bool {
	switch r {
	case SettingContentPolicyGetResponseTargetSubject, SettingContentPolicyGetResponseTargetBody:
		return true
	}
	return false
}

type SettingContentPolicyNewParams struct {
	// Identifier.
	AccountID param.Field[string]                                `path:"account_id" api:"required"`
	Enabled   param.Field[bool]                                  `json:"enabled" api:"required"`
	Name      param.Field[string]                                `json:"name" api:"required"`
	Pattern   param.Field[string]                                `json:"pattern" api:"required"`
	Targets   param.Field[[]SettingContentPolicyNewParamsTarget] `json:"targets" api:"required"`
	Notes     param.Field[string]                                `json:"notes"`
}

func (r SettingContentPolicyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The part of the email to match the pattern against.
type SettingContentPolicyNewParamsTarget string

const (
	SettingContentPolicyNewParamsTargetSubject SettingContentPolicyNewParamsTarget = "SUBJECT"
	SettingContentPolicyNewParamsTargetBody    SettingContentPolicyNewParamsTarget = "BODY"
)

func (r SettingContentPolicyNewParamsTarget) IsKnown() bool {
	switch r {
	case SettingContentPolicyNewParamsTargetSubject, SettingContentPolicyNewParamsTargetBody:
		return true
	}
	return false
}

type SettingContentPolicyNewResponseEnvelope struct {
	Errors   []SettingContentPolicyNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []SettingContentPolicyNewResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success SettingContentPolicyNewResponseEnvelopeSuccess `json:"success" api:"required"`
	// A content policy pattern that matches against the subject or body of an email.
	Result SettingContentPolicyNewResponse             `json:"result"`
	JSON   settingContentPolicyNewResponseEnvelopeJSON `json:"-"`
}

// settingContentPolicyNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingContentPolicyNewResponseEnvelope]
type settingContentPolicyNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingContentPolicyNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingContentPolicyNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingContentPolicyNewResponseEnvelopeErrors struct {
	Code             int64                                               `json:"code" api:"required"`
	Message          string                                              `json:"message" api:"required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           SettingContentPolicyNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             settingContentPolicyNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// settingContentPolicyNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SettingContentPolicyNewResponseEnvelopeErrors]
type settingContentPolicyNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingContentPolicyNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingContentPolicyNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingContentPolicyNewResponseEnvelopeErrorsSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    settingContentPolicyNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// settingContentPolicyNewResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [SettingContentPolicyNewResponseEnvelopeErrorsSource]
type settingContentPolicyNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingContentPolicyNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingContentPolicyNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SettingContentPolicyNewResponseEnvelopeMessages struct {
	Code             int64                                                 `json:"code" api:"required"`
	Message          string                                                `json:"message" api:"required"`
	DocumentationURL string                                                `json:"documentation_url"`
	Source           SettingContentPolicyNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             settingContentPolicyNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// settingContentPolicyNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingContentPolicyNewResponseEnvelopeMessages]
type settingContentPolicyNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingContentPolicyNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingContentPolicyNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingContentPolicyNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                    `json:"pointer"`
	JSON    settingContentPolicyNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// settingContentPolicyNewResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [SettingContentPolicyNewResponseEnvelopeMessagesSource]
type settingContentPolicyNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingContentPolicyNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingContentPolicyNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SettingContentPolicyNewResponseEnvelopeSuccess bool

const (
	SettingContentPolicyNewResponseEnvelopeSuccessTrue SettingContentPolicyNewResponseEnvelopeSuccess = true
)

func (r SettingContentPolicyNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingContentPolicyNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SettingContentPolicyListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// The sorting direction.
	Direction param.Field[SettingContentPolicyListParamsDirection] `query:"direction"`
	// Filter by enabled status.
	Enabled param.Field[bool] `query:"enabled"`
	// Filter by exact policy name.
	Name param.Field[string] `query:"name"`
	// Field to sort by.
	Order param.Field[SettingContentPolicyListParamsOrder] `query:"order"`
	// Current page within paginated list of results.
	Page param.Field[int64] `query:"page"`
	// The number of results per page. Maximum value is 1000.
	PerPage param.Field[int64] `query:"per_page"`
	// Search term for filtering records. Behavior may change.
	Search param.Field[string] `query:"search"`
}

// URLQuery serializes [SettingContentPolicyListParams]'s query parameters as
// `url.Values`.
func (r SettingContentPolicyListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// The sorting direction.
type SettingContentPolicyListParamsDirection string

const (
	SettingContentPolicyListParamsDirectionAsc  SettingContentPolicyListParamsDirection = "asc"
	SettingContentPolicyListParamsDirectionDesc SettingContentPolicyListParamsDirection = "desc"
)

func (r SettingContentPolicyListParamsDirection) IsKnown() bool {
	switch r {
	case SettingContentPolicyListParamsDirectionAsc, SettingContentPolicyListParamsDirectionDesc:
		return true
	}
	return false
}

// Field to sort by.
type SettingContentPolicyListParamsOrder string

const (
	SettingContentPolicyListParamsOrderName      SettingContentPolicyListParamsOrder = "name"
	SettingContentPolicyListParamsOrderCreatedAt SettingContentPolicyListParamsOrder = "created_at"
)

func (r SettingContentPolicyListParamsOrder) IsKnown() bool {
	switch r {
	case SettingContentPolicyListParamsOrderName, SettingContentPolicyListParamsOrderCreatedAt:
		return true
	}
	return false
}

type SettingContentPolicyDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type SettingContentPolicyDeleteResponseEnvelope struct {
	Errors   []SettingContentPolicyDeleteResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []SettingContentPolicyDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success SettingContentPolicyDeleteResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  SettingContentPolicyDeleteResponse                `json:"result"`
	JSON    settingContentPolicyDeleteResponseEnvelopeJSON    `json:"-"`
}

// settingContentPolicyDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingContentPolicyDeleteResponseEnvelope]
type settingContentPolicyDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingContentPolicyDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingContentPolicyDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingContentPolicyDeleteResponseEnvelopeErrors struct {
	Code             int64                                                  `json:"code" api:"required"`
	Message          string                                                 `json:"message" api:"required"`
	DocumentationURL string                                                 `json:"documentation_url"`
	Source           SettingContentPolicyDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             settingContentPolicyDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// settingContentPolicyDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingContentPolicyDeleteResponseEnvelopeErrors]
type settingContentPolicyDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingContentPolicyDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingContentPolicyDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingContentPolicyDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                     `json:"pointer"`
	JSON    settingContentPolicyDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// settingContentPolicyDeleteResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [SettingContentPolicyDeleteResponseEnvelopeErrorsSource]
type settingContentPolicyDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingContentPolicyDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingContentPolicyDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SettingContentPolicyDeleteResponseEnvelopeMessages struct {
	Code             int64                                                    `json:"code" api:"required"`
	Message          string                                                   `json:"message" api:"required"`
	DocumentationURL string                                                   `json:"documentation_url"`
	Source           SettingContentPolicyDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             settingContentPolicyDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// settingContentPolicyDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SettingContentPolicyDeleteResponseEnvelopeMessages]
type settingContentPolicyDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingContentPolicyDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingContentPolicyDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingContentPolicyDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                       `json:"pointer"`
	JSON    settingContentPolicyDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// settingContentPolicyDeleteResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [SettingContentPolicyDeleteResponseEnvelopeMessagesSource]
type settingContentPolicyDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingContentPolicyDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingContentPolicyDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SettingContentPolicyDeleteResponseEnvelopeSuccess bool

const (
	SettingContentPolicyDeleteResponseEnvelopeSuccessTrue SettingContentPolicyDeleteResponseEnvelopeSuccess = true
)

func (r SettingContentPolicyDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingContentPolicyDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SettingContentPolicyBatchParams struct {
	// Identifier.
	AccountID param.Field[string]                                  `path:"account_id" api:"required"`
	Deletes   param.Field[[]SettingContentPolicyBatchParamsDelete] `json:"deletes" api:"required"`
	Patches   param.Field[[]SettingContentPolicyBatchParamsPatch]  `json:"patches" api:"required"`
	Posts     param.Field[[]SettingContentPolicyBatchParamsPost]   `json:"posts" api:"required"`
	Puts      param.Field[[]SettingContentPolicyBatchParamsPut]    `json:"puts" api:"required"`
}

func (r SettingContentPolicyBatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingContentPolicyBatchParamsDelete struct {
	// Content policy identifier.
	ID param.Field[string] `json:"id" api:"required" format:"uuid"`
}

func (r SettingContentPolicyBatchParamsDelete) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A content policy pattern that matches against the subject or body of an email.
type SettingContentPolicyBatchParamsPatch struct {
	Enabled param.Field[bool]                                           `json:"enabled"`
	Name    param.Field[string]                                         `json:"name"`
	Notes   param.Field[string]                                         `json:"notes"`
	Pattern param.Field[string]                                         `json:"pattern"`
	Targets param.Field[[]SettingContentPolicyBatchParamsPatchesTarget] `json:"targets"`
}

func (r SettingContentPolicyBatchParamsPatch) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The part of the email to match the pattern against.
type SettingContentPolicyBatchParamsPatchesTarget string

const (
	SettingContentPolicyBatchParamsPatchesTargetSubject SettingContentPolicyBatchParamsPatchesTarget = "SUBJECT"
	SettingContentPolicyBatchParamsPatchesTargetBody    SettingContentPolicyBatchParamsPatchesTarget = "BODY"
)

func (r SettingContentPolicyBatchParamsPatchesTarget) IsKnown() bool {
	switch r {
	case SettingContentPolicyBatchParamsPatchesTargetSubject, SettingContentPolicyBatchParamsPatchesTargetBody:
		return true
	}
	return false
}

// Create a content policy.
type SettingContentPolicyBatchParamsPost struct {
	Enabled param.Field[bool]                                         `json:"enabled" api:"required"`
	Name    param.Field[string]                                       `json:"name" api:"required"`
	Pattern param.Field[string]                                       `json:"pattern" api:"required"`
	Targets param.Field[[]SettingContentPolicyBatchParamsPostsTarget] `json:"targets" api:"required"`
	Notes   param.Field[string]                                       `json:"notes"`
}

func (r SettingContentPolicyBatchParamsPost) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The part of the email to match the pattern against.
type SettingContentPolicyBatchParamsPostsTarget string

const (
	SettingContentPolicyBatchParamsPostsTargetSubject SettingContentPolicyBatchParamsPostsTarget = "SUBJECT"
	SettingContentPolicyBatchParamsPostsTargetBody    SettingContentPolicyBatchParamsPostsTarget = "BODY"
)

func (r SettingContentPolicyBatchParamsPostsTarget) IsKnown() bool {
	switch r {
	case SettingContentPolicyBatchParamsPostsTargetSubject, SettingContentPolicyBatchParamsPostsTargetBody:
		return true
	}
	return false
}

// A content policy pattern that matches against the subject or body of an email.
type SettingContentPolicyBatchParamsPut struct {
	Enabled param.Field[bool]                                        `json:"enabled" api:"required"`
	Name    param.Field[string]                                      `json:"name" api:"required"`
	Pattern param.Field[string]                                      `json:"pattern" api:"required"`
	Targets param.Field[[]SettingContentPolicyBatchParamsPutsTarget] `json:"targets" api:"required"`
	Notes   param.Field[string]                                      `json:"notes"`
}

func (r SettingContentPolicyBatchParamsPut) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The part of the email to match the pattern against.
type SettingContentPolicyBatchParamsPutsTarget string

const (
	SettingContentPolicyBatchParamsPutsTargetSubject SettingContentPolicyBatchParamsPutsTarget = "SUBJECT"
	SettingContentPolicyBatchParamsPutsTargetBody    SettingContentPolicyBatchParamsPutsTarget = "BODY"
)

func (r SettingContentPolicyBatchParamsPutsTarget) IsKnown() bool {
	switch r {
	case SettingContentPolicyBatchParamsPutsTargetSubject, SettingContentPolicyBatchParamsPutsTargetBody:
		return true
	}
	return false
}

type SettingContentPolicyBatchResponseEnvelope struct {
	Errors   []SettingContentPolicyBatchResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []SettingContentPolicyBatchResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success SettingContentPolicyBatchResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  SettingContentPolicyBatchResponse                `json:"result"`
	JSON    settingContentPolicyBatchResponseEnvelopeJSON    `json:"-"`
}

// settingContentPolicyBatchResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingContentPolicyBatchResponseEnvelope]
type settingContentPolicyBatchResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingContentPolicyBatchResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingContentPolicyBatchResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingContentPolicyBatchResponseEnvelopeErrors struct {
	Code             int64                                                 `json:"code" api:"required"`
	Message          string                                                `json:"message" api:"required"`
	DocumentationURL string                                                `json:"documentation_url"`
	Source           SettingContentPolicyBatchResponseEnvelopeErrorsSource `json:"source"`
	JSON             settingContentPolicyBatchResponseEnvelopeErrorsJSON   `json:"-"`
}

// settingContentPolicyBatchResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingContentPolicyBatchResponseEnvelopeErrors]
type settingContentPolicyBatchResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingContentPolicyBatchResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingContentPolicyBatchResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingContentPolicyBatchResponseEnvelopeErrorsSource struct {
	Pointer string                                                    `json:"pointer"`
	JSON    settingContentPolicyBatchResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// settingContentPolicyBatchResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [SettingContentPolicyBatchResponseEnvelopeErrorsSource]
type settingContentPolicyBatchResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingContentPolicyBatchResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingContentPolicyBatchResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SettingContentPolicyBatchResponseEnvelopeMessages struct {
	Code             int64                                                   `json:"code" api:"required"`
	Message          string                                                  `json:"message" api:"required"`
	DocumentationURL string                                                  `json:"documentation_url"`
	Source           SettingContentPolicyBatchResponseEnvelopeMessagesSource `json:"source"`
	JSON             settingContentPolicyBatchResponseEnvelopeMessagesJSON   `json:"-"`
}

// settingContentPolicyBatchResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingContentPolicyBatchResponseEnvelopeMessages]
type settingContentPolicyBatchResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingContentPolicyBatchResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingContentPolicyBatchResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingContentPolicyBatchResponseEnvelopeMessagesSource struct {
	Pointer string                                                      `json:"pointer"`
	JSON    settingContentPolicyBatchResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// settingContentPolicyBatchResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [SettingContentPolicyBatchResponseEnvelopeMessagesSource]
type settingContentPolicyBatchResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingContentPolicyBatchResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingContentPolicyBatchResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SettingContentPolicyBatchResponseEnvelopeSuccess bool

const (
	SettingContentPolicyBatchResponseEnvelopeSuccessTrue SettingContentPolicyBatchResponseEnvelopeSuccess = true
)

func (r SettingContentPolicyBatchResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingContentPolicyBatchResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SettingContentPolicyEditParams struct {
	// Identifier.
	AccountID param.Field[string]                                 `path:"account_id" api:"required"`
	Enabled   param.Field[bool]                                   `json:"enabled"`
	Name      param.Field[string]                                 `json:"name"`
	Notes     param.Field[string]                                 `json:"notes"`
	Pattern   param.Field[string]                                 `json:"pattern"`
	Targets   param.Field[[]SettingContentPolicyEditParamsTarget] `json:"targets"`
}

func (r SettingContentPolicyEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The part of the email to match the pattern against.
type SettingContentPolicyEditParamsTarget string

const (
	SettingContentPolicyEditParamsTargetSubject SettingContentPolicyEditParamsTarget = "SUBJECT"
	SettingContentPolicyEditParamsTargetBody    SettingContentPolicyEditParamsTarget = "BODY"
)

func (r SettingContentPolicyEditParamsTarget) IsKnown() bool {
	switch r {
	case SettingContentPolicyEditParamsTargetSubject, SettingContentPolicyEditParamsTargetBody:
		return true
	}
	return false
}

type SettingContentPolicyEditResponseEnvelope struct {
	Errors   []SettingContentPolicyEditResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []SettingContentPolicyEditResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success SettingContentPolicyEditResponseEnvelopeSuccess `json:"success" api:"required"`
	// A content policy pattern that matches against the subject or body of an email.
	Result SettingContentPolicyEditResponse             `json:"result"`
	JSON   settingContentPolicyEditResponseEnvelopeJSON `json:"-"`
}

// settingContentPolicyEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingContentPolicyEditResponseEnvelope]
type settingContentPolicyEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingContentPolicyEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingContentPolicyEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingContentPolicyEditResponseEnvelopeErrors struct {
	Code             int64                                                `json:"code" api:"required"`
	Message          string                                               `json:"message" api:"required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           SettingContentPolicyEditResponseEnvelopeErrorsSource `json:"source"`
	JSON             settingContentPolicyEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// settingContentPolicyEditResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingContentPolicyEditResponseEnvelopeErrors]
type settingContentPolicyEditResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingContentPolicyEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingContentPolicyEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingContentPolicyEditResponseEnvelopeErrorsSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    settingContentPolicyEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// settingContentPolicyEditResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [SettingContentPolicyEditResponseEnvelopeErrorsSource]
type settingContentPolicyEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingContentPolicyEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingContentPolicyEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SettingContentPolicyEditResponseEnvelopeMessages struct {
	Code             int64                                                  `json:"code" api:"required"`
	Message          string                                                 `json:"message" api:"required"`
	DocumentationURL string                                                 `json:"documentation_url"`
	Source           SettingContentPolicyEditResponseEnvelopeMessagesSource `json:"source"`
	JSON             settingContentPolicyEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// settingContentPolicyEditResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingContentPolicyEditResponseEnvelopeMessages]
type settingContentPolicyEditResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingContentPolicyEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingContentPolicyEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingContentPolicyEditResponseEnvelopeMessagesSource struct {
	Pointer string                                                     `json:"pointer"`
	JSON    settingContentPolicyEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// settingContentPolicyEditResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [SettingContentPolicyEditResponseEnvelopeMessagesSource]
type settingContentPolicyEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingContentPolicyEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingContentPolicyEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SettingContentPolicyEditResponseEnvelopeSuccess bool

const (
	SettingContentPolicyEditResponseEnvelopeSuccessTrue SettingContentPolicyEditResponseEnvelopeSuccess = true
)

func (r SettingContentPolicyEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingContentPolicyEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SettingContentPolicyGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type SettingContentPolicyGetResponseEnvelope struct {
	Errors   []SettingContentPolicyGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []SettingContentPolicyGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success SettingContentPolicyGetResponseEnvelopeSuccess `json:"success" api:"required"`
	// A content policy pattern that matches against the subject or body of an email.
	Result SettingContentPolicyGetResponse             `json:"result"`
	JSON   settingContentPolicyGetResponseEnvelopeJSON `json:"-"`
}

// settingContentPolicyGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingContentPolicyGetResponseEnvelope]
type settingContentPolicyGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingContentPolicyGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingContentPolicyGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingContentPolicyGetResponseEnvelopeErrors struct {
	Code             int64                                               `json:"code" api:"required"`
	Message          string                                              `json:"message" api:"required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           SettingContentPolicyGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             settingContentPolicyGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// settingContentPolicyGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SettingContentPolicyGetResponseEnvelopeErrors]
type settingContentPolicyGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingContentPolicyGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingContentPolicyGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingContentPolicyGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    settingContentPolicyGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// settingContentPolicyGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [SettingContentPolicyGetResponseEnvelopeErrorsSource]
type settingContentPolicyGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingContentPolicyGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingContentPolicyGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SettingContentPolicyGetResponseEnvelopeMessages struct {
	Code             int64                                                 `json:"code" api:"required"`
	Message          string                                                `json:"message" api:"required"`
	DocumentationURL string                                                `json:"documentation_url"`
	Source           SettingContentPolicyGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             settingContentPolicyGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// settingContentPolicyGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingContentPolicyGetResponseEnvelopeMessages]
type settingContentPolicyGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingContentPolicyGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingContentPolicyGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingContentPolicyGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                    `json:"pointer"`
	JSON    settingContentPolicyGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// settingContentPolicyGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [SettingContentPolicyGetResponseEnvelopeMessagesSource]
type settingContentPolicyGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingContentPolicyGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingContentPolicyGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SettingContentPolicyGetResponseEnvelopeSuccess bool

const (
	SettingContentPolicyGetResponseEnvelopeSuccessTrue SettingContentPolicyGetResponseEnvelopeSuccess = true
)

func (r SettingContentPolicyGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingContentPolicyGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
