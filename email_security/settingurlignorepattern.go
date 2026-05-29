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

// SettingURLIgnorePatternService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingURLIgnorePatternService] method instead.
type SettingURLIgnorePatternService struct {
	Options []option.RequestOption
}

// NewSettingURLIgnorePatternService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingURLIgnorePatternService(opts ...option.RequestOption) (r *SettingURLIgnorePatternService) {
	r = &SettingURLIgnorePatternService{}
	r.Options = opts
	return
}

// Creates a new URL rewrite ignore pattern. URLs matching this pattern will not be
// rewritten.
func (r *SettingURLIgnorePatternService) New(ctx context.Context, params SettingURLIgnorePatternNewParams, opts ...option.RequestOption) (res *SettingURLIgnorePatternNewResponse, err error) {
	var env SettingURLIgnorePatternNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/url_ignore_patterns", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns a paginated list of URL rewrite ignore patterns for the account. URLs
// matching these patterns will not be rewritten.
func (r *SettingURLIgnorePatternService) List(ctx context.Context, params SettingURLIgnorePatternListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[SettingURLIgnorePatternListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/url_ignore_patterns", params.AccountID)
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

// Returns a paginated list of URL rewrite ignore patterns for the account. URLs
// matching these patterns will not be rewritten.
func (r *SettingURLIgnorePatternService) ListAutoPaging(ctx context.Context, params SettingURLIgnorePatternListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[SettingURLIgnorePatternListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Removes a URL rewrite ignore pattern. After deletion, URLs matching this pattern
// will be rewritten again.
func (r *SettingURLIgnorePatternService) Delete(ctx context.Context, patternID string, body SettingURLIgnorePatternDeleteParams, opts ...option.RequestOption) (res *SettingURLIgnorePatternDeleteResponse, err error) {
	var env SettingURLIgnorePatternDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if patternID == "" {
		err = errors.New("missing required pattern_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/url_ignore_patterns/%s", body.AccountID, patternID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Updates an existing URL rewrite ignore pattern. Only provided fields will be
// modified.
func (r *SettingURLIgnorePatternService) Edit(ctx context.Context, patternID string, params SettingURLIgnorePatternEditParams, opts ...option.RequestOption) (res *SettingURLIgnorePatternEditResponse, err error) {
	var env SettingURLIgnorePatternEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if patternID == "" {
		err = errors.New("missing required pattern_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/url_ignore_patterns/%s", params.AccountID, patternID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns a single URL rewrite ignore pattern by its identifier.
func (r *SettingURLIgnorePatternService) Get(ctx context.Context, patternID string, query SettingURLIgnorePatternGetParams, opts ...option.RequestOption) (res *SettingURLIgnorePatternGetResponse, err error) {
	var env SettingURLIgnorePatternGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if patternID == "" {
		err = errors.New("missing required pattern_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/url_ignore_patterns/%s", query.AccountID, patternID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// A URL ignore pattern that exempts matching URLs from being rewritten by Email
// Security.
type SettingURLIgnorePatternNewResponse struct {
	// URL ignore pattern identifier
	ID        string    `json:"id" api:"required" format:"uuid"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Regular expression matching URLs that should not be rewritten.
	Pattern string `json:"pattern" api:"required"`
	// Optional note describing the reason for the ignore pattern.
	Comments string `json:"comments" api:"nullable"`
	// Deprecated, use `modified_at` instead. End of life: November 1, 2026.
	//
	// Deprecated: deprecated
	LastModified time.Time                              `json:"last_modified" format:"date-time"`
	ModifiedAt   time.Time                              `json:"modified_at" format:"date-time"`
	JSON         settingURLIgnorePatternNewResponseJSON `json:"-"`
}

// settingURLIgnorePatternNewResponseJSON contains the JSON metadata for the struct
// [SettingURLIgnorePatternNewResponse]
type settingURLIgnorePatternNewResponseJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	Pattern      apijson.Field
	Comments     apijson.Field
	LastModified apijson.Field
	ModifiedAt   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SettingURLIgnorePatternNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingURLIgnorePatternNewResponseJSON) RawJSON() string {
	return r.raw
}

// A URL ignore pattern that exempts matching URLs from being rewritten by Email
// Security.
type SettingURLIgnorePatternListResponse struct {
	// URL ignore pattern identifier
	ID        string    `json:"id" api:"required" format:"uuid"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Regular expression matching URLs that should not be rewritten.
	Pattern string `json:"pattern" api:"required"`
	// Optional note describing the reason for the ignore pattern.
	Comments string `json:"comments" api:"nullable"`
	// Deprecated, use `modified_at` instead. End of life: November 1, 2026.
	//
	// Deprecated: deprecated
	LastModified time.Time                               `json:"last_modified" format:"date-time"`
	ModifiedAt   time.Time                               `json:"modified_at" format:"date-time"`
	JSON         settingURLIgnorePatternListResponseJSON `json:"-"`
}

// settingURLIgnorePatternListResponseJSON contains the JSON metadata for the
// struct [SettingURLIgnorePatternListResponse]
type settingURLIgnorePatternListResponseJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	Pattern      apijson.Field
	Comments     apijson.Field
	LastModified apijson.Field
	ModifiedAt   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SettingURLIgnorePatternListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingURLIgnorePatternListResponseJSON) RawJSON() string {
	return r.raw
}

type SettingURLIgnorePatternDeleteResponse struct {
	// URL ignore pattern identifier
	ID   string                                    `json:"id" api:"required" format:"uuid"`
	JSON settingURLIgnorePatternDeleteResponseJSON `json:"-"`
}

// settingURLIgnorePatternDeleteResponseJSON contains the JSON metadata for the
// struct [SettingURLIgnorePatternDeleteResponse]
type settingURLIgnorePatternDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingURLIgnorePatternDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingURLIgnorePatternDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// A URL ignore pattern that exempts matching URLs from being rewritten by Email
// Security.
type SettingURLIgnorePatternEditResponse struct {
	// URL ignore pattern identifier
	ID        string    `json:"id" api:"required" format:"uuid"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Regular expression matching URLs that should not be rewritten.
	Pattern string `json:"pattern" api:"required"`
	// Optional note describing the reason for the ignore pattern.
	Comments string `json:"comments" api:"nullable"`
	// Deprecated, use `modified_at` instead. End of life: November 1, 2026.
	//
	// Deprecated: deprecated
	LastModified time.Time                               `json:"last_modified" format:"date-time"`
	ModifiedAt   time.Time                               `json:"modified_at" format:"date-time"`
	JSON         settingURLIgnorePatternEditResponseJSON `json:"-"`
}

// settingURLIgnorePatternEditResponseJSON contains the JSON metadata for the
// struct [SettingURLIgnorePatternEditResponse]
type settingURLIgnorePatternEditResponseJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	Pattern      apijson.Field
	Comments     apijson.Field
	LastModified apijson.Field
	ModifiedAt   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SettingURLIgnorePatternEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingURLIgnorePatternEditResponseJSON) RawJSON() string {
	return r.raw
}

// A URL ignore pattern that exempts matching URLs from being rewritten by Email
// Security.
type SettingURLIgnorePatternGetResponse struct {
	// URL ignore pattern identifier
	ID        string    `json:"id" api:"required" format:"uuid"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Regular expression matching URLs that should not be rewritten.
	Pattern string `json:"pattern" api:"required"`
	// Optional note describing the reason for the ignore pattern.
	Comments string `json:"comments" api:"nullable"`
	// Deprecated, use `modified_at` instead. End of life: November 1, 2026.
	//
	// Deprecated: deprecated
	LastModified time.Time                              `json:"last_modified" format:"date-time"`
	ModifiedAt   time.Time                              `json:"modified_at" format:"date-time"`
	JSON         settingURLIgnorePatternGetResponseJSON `json:"-"`
}

// settingURLIgnorePatternGetResponseJSON contains the JSON metadata for the struct
// [SettingURLIgnorePatternGetResponse]
type settingURLIgnorePatternGetResponseJSON struct {
	ID           apijson.Field
	CreatedAt    apijson.Field
	Pattern      apijson.Field
	Comments     apijson.Field
	LastModified apijson.Field
	ModifiedAt   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SettingURLIgnorePatternGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingURLIgnorePatternGetResponseJSON) RawJSON() string {
	return r.raw
}

type SettingURLIgnorePatternNewParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Regular expression matching URLs that should not be rewritten.
	Pattern param.Field[string] `json:"pattern" api:"required"`
	// Optional note describing the reason for the ignore pattern.
	Comments param.Field[string] `json:"comments"`
}

func (r SettingURLIgnorePatternNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingURLIgnorePatternNewResponseEnvelope struct {
	Errors   []SettingURLIgnorePatternNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []SettingURLIgnorePatternNewResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success SettingURLIgnorePatternNewResponseEnvelopeSuccess `json:"success" api:"required"`
	// A URL ignore pattern that exempts matching URLs from being rewritten by Email
	// Security.
	Result SettingURLIgnorePatternNewResponse             `json:"result"`
	JSON   settingURLIgnorePatternNewResponseEnvelopeJSON `json:"-"`
}

// settingURLIgnorePatternNewResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingURLIgnorePatternNewResponseEnvelope]
type settingURLIgnorePatternNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingURLIgnorePatternNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingURLIgnorePatternNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingURLIgnorePatternNewResponseEnvelopeErrors struct {
	Code             int64                                                  `json:"code" api:"required"`
	Message          string                                                 `json:"message" api:"required"`
	DocumentationURL string                                                 `json:"documentation_url"`
	Source           SettingURLIgnorePatternNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             settingURLIgnorePatternNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// settingURLIgnorePatternNewResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingURLIgnorePatternNewResponseEnvelopeErrors]
type settingURLIgnorePatternNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingURLIgnorePatternNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingURLIgnorePatternNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingURLIgnorePatternNewResponseEnvelopeErrorsSource struct {
	Pointer string                                                     `json:"pointer"`
	JSON    settingURLIgnorePatternNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// settingURLIgnorePatternNewResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [SettingURLIgnorePatternNewResponseEnvelopeErrorsSource]
type settingURLIgnorePatternNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingURLIgnorePatternNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingURLIgnorePatternNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SettingURLIgnorePatternNewResponseEnvelopeMessages struct {
	Code             int64                                                    `json:"code" api:"required"`
	Message          string                                                   `json:"message" api:"required"`
	DocumentationURL string                                                   `json:"documentation_url"`
	Source           SettingURLIgnorePatternNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             settingURLIgnorePatternNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// settingURLIgnorePatternNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SettingURLIgnorePatternNewResponseEnvelopeMessages]
type settingURLIgnorePatternNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingURLIgnorePatternNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingURLIgnorePatternNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingURLIgnorePatternNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                       `json:"pointer"`
	JSON    settingURLIgnorePatternNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// settingURLIgnorePatternNewResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [SettingURLIgnorePatternNewResponseEnvelopeMessagesSource]
type settingURLIgnorePatternNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingURLIgnorePatternNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingURLIgnorePatternNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SettingURLIgnorePatternNewResponseEnvelopeSuccess bool

const (
	SettingURLIgnorePatternNewResponseEnvelopeSuccessTrue SettingURLIgnorePatternNewResponseEnvelopeSuccess = true
)

func (r SettingURLIgnorePatternNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingURLIgnorePatternNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SettingURLIgnorePatternListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Current page within paginated list of results.
	Page param.Field[int64] `query:"page"`
	// The number of results per page. Maximum value is 1000.
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [SettingURLIgnorePatternListParams]'s query parameters as
// `url.Values`.
func (r SettingURLIgnorePatternListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type SettingURLIgnorePatternDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type SettingURLIgnorePatternDeleteResponseEnvelope struct {
	Errors   []SettingURLIgnorePatternDeleteResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []SettingURLIgnorePatternDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success SettingURLIgnorePatternDeleteResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  SettingURLIgnorePatternDeleteResponse                `json:"result"`
	JSON    settingURLIgnorePatternDeleteResponseEnvelopeJSON    `json:"-"`
}

// settingURLIgnorePatternDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingURLIgnorePatternDeleteResponseEnvelope]
type settingURLIgnorePatternDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingURLIgnorePatternDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingURLIgnorePatternDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingURLIgnorePatternDeleteResponseEnvelopeErrors struct {
	Code             int64                                                     `json:"code" api:"required"`
	Message          string                                                    `json:"message" api:"required"`
	DocumentationURL string                                                    `json:"documentation_url"`
	Source           SettingURLIgnorePatternDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             settingURLIgnorePatternDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// settingURLIgnorePatternDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [SettingURLIgnorePatternDeleteResponseEnvelopeErrors]
type settingURLIgnorePatternDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingURLIgnorePatternDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingURLIgnorePatternDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingURLIgnorePatternDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                        `json:"pointer"`
	JSON    settingURLIgnorePatternDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// settingURLIgnorePatternDeleteResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [SettingURLIgnorePatternDeleteResponseEnvelopeErrorsSource]
type settingURLIgnorePatternDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingURLIgnorePatternDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingURLIgnorePatternDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SettingURLIgnorePatternDeleteResponseEnvelopeMessages struct {
	Code             int64                                                       `json:"code" api:"required"`
	Message          string                                                      `json:"message" api:"required"`
	DocumentationURL string                                                      `json:"documentation_url"`
	Source           SettingURLIgnorePatternDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             settingURLIgnorePatternDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// settingURLIgnorePatternDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SettingURLIgnorePatternDeleteResponseEnvelopeMessages]
type settingURLIgnorePatternDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingURLIgnorePatternDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingURLIgnorePatternDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingURLIgnorePatternDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                          `json:"pointer"`
	JSON    settingURLIgnorePatternDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// settingURLIgnorePatternDeleteResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [SettingURLIgnorePatternDeleteResponseEnvelopeMessagesSource]
type settingURLIgnorePatternDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingURLIgnorePatternDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingURLIgnorePatternDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SettingURLIgnorePatternDeleteResponseEnvelopeSuccess bool

const (
	SettingURLIgnorePatternDeleteResponseEnvelopeSuccessTrue SettingURLIgnorePatternDeleteResponseEnvelopeSuccess = true
)

func (r SettingURLIgnorePatternDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingURLIgnorePatternDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SettingURLIgnorePatternEditParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Optional note describing the reason for the ignore pattern.
	Comments param.Field[string] `json:"comments"`
	// Regular expression matching URLs that should not be rewritten.
	Pattern param.Field[string] `json:"pattern"`
}

func (r SettingURLIgnorePatternEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingURLIgnorePatternEditResponseEnvelope struct {
	Errors   []SettingURLIgnorePatternEditResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []SettingURLIgnorePatternEditResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success SettingURLIgnorePatternEditResponseEnvelopeSuccess `json:"success" api:"required"`
	// A URL ignore pattern that exempts matching URLs from being rewritten by Email
	// Security.
	Result SettingURLIgnorePatternEditResponse             `json:"result"`
	JSON   settingURLIgnorePatternEditResponseEnvelopeJSON `json:"-"`
}

// settingURLIgnorePatternEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingURLIgnorePatternEditResponseEnvelope]
type settingURLIgnorePatternEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingURLIgnorePatternEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingURLIgnorePatternEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingURLIgnorePatternEditResponseEnvelopeErrors struct {
	Code             int64                                                   `json:"code" api:"required"`
	Message          string                                                  `json:"message" api:"required"`
	DocumentationURL string                                                  `json:"documentation_url"`
	Source           SettingURLIgnorePatternEditResponseEnvelopeErrorsSource `json:"source"`
	JSON             settingURLIgnorePatternEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// settingURLIgnorePatternEditResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingURLIgnorePatternEditResponseEnvelopeErrors]
type settingURLIgnorePatternEditResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingURLIgnorePatternEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingURLIgnorePatternEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingURLIgnorePatternEditResponseEnvelopeErrorsSource struct {
	Pointer string                                                      `json:"pointer"`
	JSON    settingURLIgnorePatternEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// settingURLIgnorePatternEditResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [SettingURLIgnorePatternEditResponseEnvelopeErrorsSource]
type settingURLIgnorePatternEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingURLIgnorePatternEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingURLIgnorePatternEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SettingURLIgnorePatternEditResponseEnvelopeMessages struct {
	Code             int64                                                     `json:"code" api:"required"`
	Message          string                                                    `json:"message" api:"required"`
	DocumentationURL string                                                    `json:"documentation_url"`
	Source           SettingURLIgnorePatternEditResponseEnvelopeMessagesSource `json:"source"`
	JSON             settingURLIgnorePatternEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// settingURLIgnorePatternEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SettingURLIgnorePatternEditResponseEnvelopeMessages]
type settingURLIgnorePatternEditResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingURLIgnorePatternEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingURLIgnorePatternEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingURLIgnorePatternEditResponseEnvelopeMessagesSource struct {
	Pointer string                                                        `json:"pointer"`
	JSON    settingURLIgnorePatternEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// settingURLIgnorePatternEditResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [SettingURLIgnorePatternEditResponseEnvelopeMessagesSource]
type settingURLIgnorePatternEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingURLIgnorePatternEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingURLIgnorePatternEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SettingURLIgnorePatternEditResponseEnvelopeSuccess bool

const (
	SettingURLIgnorePatternEditResponseEnvelopeSuccessTrue SettingURLIgnorePatternEditResponseEnvelopeSuccess = true
)

func (r SettingURLIgnorePatternEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingURLIgnorePatternEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SettingURLIgnorePatternGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type SettingURLIgnorePatternGetResponseEnvelope struct {
	Errors   []SettingURLIgnorePatternGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []SettingURLIgnorePatternGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success SettingURLIgnorePatternGetResponseEnvelopeSuccess `json:"success" api:"required"`
	// A URL ignore pattern that exempts matching URLs from being rewritten by Email
	// Security.
	Result SettingURLIgnorePatternGetResponse             `json:"result"`
	JSON   settingURLIgnorePatternGetResponseEnvelopeJSON `json:"-"`
}

// settingURLIgnorePatternGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [SettingURLIgnorePatternGetResponseEnvelope]
type settingURLIgnorePatternGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingURLIgnorePatternGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingURLIgnorePatternGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingURLIgnorePatternGetResponseEnvelopeErrors struct {
	Code             int64                                                  `json:"code" api:"required"`
	Message          string                                                 `json:"message" api:"required"`
	DocumentationURL string                                                 `json:"documentation_url"`
	Source           SettingURLIgnorePatternGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             settingURLIgnorePatternGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// settingURLIgnorePatternGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingURLIgnorePatternGetResponseEnvelopeErrors]
type settingURLIgnorePatternGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingURLIgnorePatternGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingURLIgnorePatternGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingURLIgnorePatternGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                     `json:"pointer"`
	JSON    settingURLIgnorePatternGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// settingURLIgnorePatternGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [SettingURLIgnorePatternGetResponseEnvelopeErrorsSource]
type settingURLIgnorePatternGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingURLIgnorePatternGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingURLIgnorePatternGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SettingURLIgnorePatternGetResponseEnvelopeMessages struct {
	Code             int64                                                    `json:"code" api:"required"`
	Message          string                                                   `json:"message" api:"required"`
	DocumentationURL string                                                   `json:"documentation_url"`
	Source           SettingURLIgnorePatternGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             settingURLIgnorePatternGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// settingURLIgnorePatternGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SettingURLIgnorePatternGetResponseEnvelopeMessages]
type settingURLIgnorePatternGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingURLIgnorePatternGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingURLIgnorePatternGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingURLIgnorePatternGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                       `json:"pointer"`
	JSON    settingURLIgnorePatternGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// settingURLIgnorePatternGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [SettingURLIgnorePatternGetResponseEnvelopeMessagesSource]
type settingURLIgnorePatternGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingURLIgnorePatternGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingURLIgnorePatternGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SettingURLIgnorePatternGetResponseEnvelopeSuccess bool

const (
	SettingURLIgnorePatternGetResponseEnvelopeSuccessTrue SettingURLIgnorePatternGetResponseEnvelopeSuccess = true
)

func (r SettingURLIgnorePatternGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingURLIgnorePatternGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
