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

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
)

// SettingBlockSenderService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingBlockSenderService] method instead.
type SettingBlockSenderService struct {
	Options []option.RequestOption
}

// NewSettingBlockSenderService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingBlockSenderService(opts ...option.RequestOption) (r *SettingBlockSenderService) {
	r = &SettingBlockSenderService{}
	r.Options = opts
	return
}

// Creates a new blocked sender pattern. Emails matching this pattern will be
// blocked from delivery. Patterns can be email addresses, domains, or IP
// addresses, and support regular expressions.
func (r *SettingBlockSenderService) New(ctx context.Context, params SettingBlockSenderNewParams, opts ...option.RequestOption) (res *SettingBlockSenderNewResponse, err error) {
	var env SettingBlockSenderNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/block_senders", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns a paginated list of blocked email sender patterns. These patterns
// prevent emails from matching senders from being delivered. Supports filtering by
// pattern type and searching across patterns.
func (r *SettingBlockSenderService) List(ctx context.Context, params SettingBlockSenderListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[SettingBlockSenderListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/block_senders", params.AccountID)
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

// Returns a paginated list of blocked email sender patterns. These patterns
// prevent emails from matching senders from being delivered. Supports filtering by
// pattern type and searching across patterns.
func (r *SettingBlockSenderService) ListAutoPaging(ctx context.Context, params SettingBlockSenderListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[SettingBlockSenderListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Removes a blocked sender pattern. After deletion, emails from this sender will
// no longer be automatically blocked based on this rule.
func (r *SettingBlockSenderService) Delete(ctx context.Context, patternID string, body SettingBlockSenderDeleteParams, opts ...option.RequestOption) (res *SettingBlockSenderDeleteResponse, err error) {
	var env SettingBlockSenderDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if patternID == "" {
		err = errors.New("missing required pattern_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/block_senders/%s", body.AccountID, patternID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Updates an existing blocked sender pattern. Only provided fields will be
// modified. The pattern will continue blocking emails until deleted.
func (r *SettingBlockSenderService) Edit(ctx context.Context, patternID string, params SettingBlockSenderEditParams, opts ...option.RequestOption) (res *SettingBlockSenderEditResponse, err error) {
	var env SettingBlockSenderEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if patternID == "" {
		err = errors.New("missing required pattern_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/block_senders/%s", params.AccountID, patternID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieves details for a specific blocked sender pattern including its pattern
// type, value, and metadata.
func (r *SettingBlockSenderService) Get(ctx context.Context, patternID string, query SettingBlockSenderGetParams, opts ...option.RequestOption) (res *SettingBlockSenderGetResponse, err error) {
	var env SettingBlockSenderGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if patternID == "" {
		err = errors.New("missing required pattern_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/block_senders/%s", query.AccountID, patternID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// A blocked sender pattern
type SettingBlockSenderNewResponse struct {
	// Blocked sender pattern identifier
	ID        string    `json:"id" format:"uuid"`
	Comments  string    `json:"comments" api:"nullable"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	IsRegex   bool      `json:"is_regex"`
	// Deprecated, use `modified_at` instead. End of life: November 1, 2026.
	//
	// Deprecated: deprecated
	LastModified time.Time `json:"last_modified" format:"date-time"`
	ModifiedAt   time.Time `json:"modified_at" format:"date-time"`
	Pattern      string    `json:"pattern"`
	// Type of pattern matching. Note: UNKNOWN is deprecated and cannot be used when
	// creating or updating policies, but may be returned for existing entries.
	PatternType SettingBlockSenderNewResponsePatternType `json:"pattern_type"`
	JSON        settingBlockSenderNewResponseJSON        `json:"-"`
}

// settingBlockSenderNewResponseJSON contains the JSON metadata for the struct
// [SettingBlockSenderNewResponse]
type settingBlockSenderNewResponseJSON struct {
	ID           apijson.Field
	Comments     apijson.Field
	CreatedAt    apijson.Field
	IsRegex      apijson.Field
	LastModified apijson.Field
	ModifiedAt   apijson.Field
	Pattern      apijson.Field
	PatternType  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SettingBlockSenderNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBlockSenderNewResponseJSON) RawJSON() string {
	return r.raw
}

// Type of pattern matching. Note: UNKNOWN is deprecated and cannot be used when
// creating or updating policies, but may be returned for existing entries.
type SettingBlockSenderNewResponsePatternType string

const (
	SettingBlockSenderNewResponsePatternTypeEmail   SettingBlockSenderNewResponsePatternType = "EMAIL"
	SettingBlockSenderNewResponsePatternTypeDomain  SettingBlockSenderNewResponsePatternType = "DOMAIN"
	SettingBlockSenderNewResponsePatternTypeIP      SettingBlockSenderNewResponsePatternType = "IP"
	SettingBlockSenderNewResponsePatternTypeUnknown SettingBlockSenderNewResponsePatternType = "UNKNOWN"
)

func (r SettingBlockSenderNewResponsePatternType) IsKnown() bool {
	switch r {
	case SettingBlockSenderNewResponsePatternTypeEmail, SettingBlockSenderNewResponsePatternTypeDomain, SettingBlockSenderNewResponsePatternTypeIP, SettingBlockSenderNewResponsePatternTypeUnknown:
		return true
	}
	return false
}

// A blocked sender pattern
type SettingBlockSenderListResponse struct {
	// Blocked sender pattern identifier
	ID        string    `json:"id" format:"uuid"`
	Comments  string    `json:"comments" api:"nullable"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	IsRegex   bool      `json:"is_regex"`
	// Deprecated, use `modified_at` instead. End of life: November 1, 2026.
	//
	// Deprecated: deprecated
	LastModified time.Time `json:"last_modified" format:"date-time"`
	ModifiedAt   time.Time `json:"modified_at" format:"date-time"`
	Pattern      string    `json:"pattern"`
	// Type of pattern matching. Note: UNKNOWN is deprecated and cannot be used when
	// creating or updating policies, but may be returned for existing entries.
	PatternType SettingBlockSenderListResponsePatternType `json:"pattern_type"`
	JSON        settingBlockSenderListResponseJSON        `json:"-"`
}

// settingBlockSenderListResponseJSON contains the JSON metadata for the struct
// [SettingBlockSenderListResponse]
type settingBlockSenderListResponseJSON struct {
	ID           apijson.Field
	Comments     apijson.Field
	CreatedAt    apijson.Field
	IsRegex      apijson.Field
	LastModified apijson.Field
	ModifiedAt   apijson.Field
	Pattern      apijson.Field
	PatternType  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SettingBlockSenderListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBlockSenderListResponseJSON) RawJSON() string {
	return r.raw
}

// Type of pattern matching. Note: UNKNOWN is deprecated and cannot be used when
// creating or updating policies, but may be returned for existing entries.
type SettingBlockSenderListResponsePatternType string

const (
	SettingBlockSenderListResponsePatternTypeEmail   SettingBlockSenderListResponsePatternType = "EMAIL"
	SettingBlockSenderListResponsePatternTypeDomain  SettingBlockSenderListResponsePatternType = "DOMAIN"
	SettingBlockSenderListResponsePatternTypeIP      SettingBlockSenderListResponsePatternType = "IP"
	SettingBlockSenderListResponsePatternTypeUnknown SettingBlockSenderListResponsePatternType = "UNKNOWN"
)

func (r SettingBlockSenderListResponsePatternType) IsKnown() bool {
	switch r {
	case SettingBlockSenderListResponsePatternTypeEmail, SettingBlockSenderListResponsePatternTypeDomain, SettingBlockSenderListResponsePatternTypeIP, SettingBlockSenderListResponsePatternTypeUnknown:
		return true
	}
	return false
}

type SettingBlockSenderDeleteResponse struct {
	// Blocked sender pattern identifier
	ID   string                               `json:"id" api:"required" format:"uuid"`
	JSON settingBlockSenderDeleteResponseJSON `json:"-"`
}

// settingBlockSenderDeleteResponseJSON contains the JSON metadata for the struct
// [SettingBlockSenderDeleteResponse]
type settingBlockSenderDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingBlockSenderDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBlockSenderDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// A blocked sender pattern
type SettingBlockSenderEditResponse struct {
	// Blocked sender pattern identifier
	ID        string    `json:"id" format:"uuid"`
	Comments  string    `json:"comments" api:"nullable"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	IsRegex   bool      `json:"is_regex"`
	// Deprecated, use `modified_at` instead. End of life: November 1, 2026.
	//
	// Deprecated: deprecated
	LastModified time.Time `json:"last_modified" format:"date-time"`
	ModifiedAt   time.Time `json:"modified_at" format:"date-time"`
	Pattern      string    `json:"pattern"`
	// Type of pattern matching. Note: UNKNOWN is deprecated and cannot be used when
	// creating or updating policies, but may be returned for existing entries.
	PatternType SettingBlockSenderEditResponsePatternType `json:"pattern_type"`
	JSON        settingBlockSenderEditResponseJSON        `json:"-"`
}

// settingBlockSenderEditResponseJSON contains the JSON metadata for the struct
// [SettingBlockSenderEditResponse]
type settingBlockSenderEditResponseJSON struct {
	ID           apijson.Field
	Comments     apijson.Field
	CreatedAt    apijson.Field
	IsRegex      apijson.Field
	LastModified apijson.Field
	ModifiedAt   apijson.Field
	Pattern      apijson.Field
	PatternType  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SettingBlockSenderEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBlockSenderEditResponseJSON) RawJSON() string {
	return r.raw
}

// Type of pattern matching. Note: UNKNOWN is deprecated and cannot be used when
// creating or updating policies, but may be returned for existing entries.
type SettingBlockSenderEditResponsePatternType string

const (
	SettingBlockSenderEditResponsePatternTypeEmail   SettingBlockSenderEditResponsePatternType = "EMAIL"
	SettingBlockSenderEditResponsePatternTypeDomain  SettingBlockSenderEditResponsePatternType = "DOMAIN"
	SettingBlockSenderEditResponsePatternTypeIP      SettingBlockSenderEditResponsePatternType = "IP"
	SettingBlockSenderEditResponsePatternTypeUnknown SettingBlockSenderEditResponsePatternType = "UNKNOWN"
)

func (r SettingBlockSenderEditResponsePatternType) IsKnown() bool {
	switch r {
	case SettingBlockSenderEditResponsePatternTypeEmail, SettingBlockSenderEditResponsePatternTypeDomain, SettingBlockSenderEditResponsePatternTypeIP, SettingBlockSenderEditResponsePatternTypeUnknown:
		return true
	}
	return false
}

// A blocked sender pattern
type SettingBlockSenderGetResponse struct {
	// Blocked sender pattern identifier
	ID        string    `json:"id" format:"uuid"`
	Comments  string    `json:"comments" api:"nullable"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	IsRegex   bool      `json:"is_regex"`
	// Deprecated, use `modified_at` instead. End of life: November 1, 2026.
	//
	// Deprecated: deprecated
	LastModified time.Time `json:"last_modified" format:"date-time"`
	ModifiedAt   time.Time `json:"modified_at" format:"date-time"`
	Pattern      string    `json:"pattern"`
	// Type of pattern matching. Note: UNKNOWN is deprecated and cannot be used when
	// creating or updating policies, but may be returned for existing entries.
	PatternType SettingBlockSenderGetResponsePatternType `json:"pattern_type"`
	JSON        settingBlockSenderGetResponseJSON        `json:"-"`
}

// settingBlockSenderGetResponseJSON contains the JSON metadata for the struct
// [SettingBlockSenderGetResponse]
type settingBlockSenderGetResponseJSON struct {
	ID           apijson.Field
	Comments     apijson.Field
	CreatedAt    apijson.Field
	IsRegex      apijson.Field
	LastModified apijson.Field
	ModifiedAt   apijson.Field
	Pattern      apijson.Field
	PatternType  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SettingBlockSenderGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBlockSenderGetResponseJSON) RawJSON() string {
	return r.raw
}

// Type of pattern matching. Note: UNKNOWN is deprecated and cannot be used when
// creating or updating policies, but may be returned for existing entries.
type SettingBlockSenderGetResponsePatternType string

const (
	SettingBlockSenderGetResponsePatternTypeEmail   SettingBlockSenderGetResponsePatternType = "EMAIL"
	SettingBlockSenderGetResponsePatternTypeDomain  SettingBlockSenderGetResponsePatternType = "DOMAIN"
	SettingBlockSenderGetResponsePatternTypeIP      SettingBlockSenderGetResponsePatternType = "IP"
	SettingBlockSenderGetResponsePatternTypeUnknown SettingBlockSenderGetResponsePatternType = "UNKNOWN"
)

func (r SettingBlockSenderGetResponsePatternType) IsKnown() bool {
	switch r {
	case SettingBlockSenderGetResponsePatternTypeEmail, SettingBlockSenderGetResponsePatternTypeDomain, SettingBlockSenderGetResponsePatternTypeIP, SettingBlockSenderGetResponsePatternTypeUnknown:
		return true
	}
	return false
}

type SettingBlockSenderNewParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	IsRegex   param.Field[bool]   `json:"is_regex" api:"required"`
	Pattern   param.Field[string] `json:"pattern" api:"required"`
	// Type of pattern matching. Note: UNKNOWN is deprecated and cannot be used when
	// creating or updating policies, but may be returned for existing entries.
	PatternType param.Field[SettingBlockSenderNewParamsPatternType] `json:"pattern_type" api:"required"`
	Comments    param.Field[string]                                 `json:"comments"`
}

func (r SettingBlockSenderNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of pattern matching. Note: UNKNOWN is deprecated and cannot be used when
// creating or updating policies, but may be returned for existing entries.
type SettingBlockSenderNewParamsPatternType string

const (
	SettingBlockSenderNewParamsPatternTypeEmail   SettingBlockSenderNewParamsPatternType = "EMAIL"
	SettingBlockSenderNewParamsPatternTypeDomain  SettingBlockSenderNewParamsPatternType = "DOMAIN"
	SettingBlockSenderNewParamsPatternTypeIP      SettingBlockSenderNewParamsPatternType = "IP"
	SettingBlockSenderNewParamsPatternTypeUnknown SettingBlockSenderNewParamsPatternType = "UNKNOWN"
)

func (r SettingBlockSenderNewParamsPatternType) IsKnown() bool {
	switch r {
	case SettingBlockSenderNewParamsPatternTypeEmail, SettingBlockSenderNewParamsPatternTypeDomain, SettingBlockSenderNewParamsPatternTypeIP, SettingBlockSenderNewParamsPatternTypeUnknown:
		return true
	}
	return false
}

type SettingBlockSenderNewResponseEnvelope struct {
	Errors   []SettingBlockSenderNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []SettingBlockSenderNewResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success SettingBlockSenderNewResponseEnvelopeSuccess `json:"success" api:"required"`
	// A blocked sender pattern
	Result SettingBlockSenderNewResponse             `json:"result"`
	JSON   settingBlockSenderNewResponseEnvelopeJSON `json:"-"`
}

// settingBlockSenderNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingBlockSenderNewResponseEnvelope]
type settingBlockSenderNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingBlockSenderNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBlockSenderNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingBlockSenderNewResponseEnvelopeErrors struct {
	Code             int64                                             `json:"code" api:"required"`
	Message          string                                            `json:"message" api:"required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           SettingBlockSenderNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             settingBlockSenderNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// settingBlockSenderNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SettingBlockSenderNewResponseEnvelopeErrors]
type settingBlockSenderNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingBlockSenderNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBlockSenderNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingBlockSenderNewResponseEnvelopeErrorsSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    settingBlockSenderNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// settingBlockSenderNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [SettingBlockSenderNewResponseEnvelopeErrorsSource]
type settingBlockSenderNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingBlockSenderNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBlockSenderNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SettingBlockSenderNewResponseEnvelopeMessages struct {
	Code             int64                                               `json:"code" api:"required"`
	Message          string                                              `json:"message" api:"required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           SettingBlockSenderNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             settingBlockSenderNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// settingBlockSenderNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SettingBlockSenderNewResponseEnvelopeMessages]
type settingBlockSenderNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingBlockSenderNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBlockSenderNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingBlockSenderNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    settingBlockSenderNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// settingBlockSenderNewResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [SettingBlockSenderNewResponseEnvelopeMessagesSource]
type settingBlockSenderNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingBlockSenderNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBlockSenderNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SettingBlockSenderNewResponseEnvelopeSuccess bool

const (
	SettingBlockSenderNewResponseEnvelopeSuccessTrue SettingBlockSenderNewResponseEnvelopeSuccess = true
)

func (r SettingBlockSenderNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingBlockSenderNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SettingBlockSenderListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// The sorting direction.
	Direction param.Field[SettingBlockSenderListParamsDirection] `query:"direction"`
	// Field to sort by.
	Order param.Field[SettingBlockSenderListParamsOrder] `query:"order"`
	// Current page within paginated list of results.
	Page param.Field[int64] `query:"page"`
	// Filter by pattern value.
	Pattern param.Field[string] `query:"pattern"`
	// Filter by pattern type.
	PatternType param.Field[SettingBlockSenderListParamsPatternType] `query:"pattern_type"`
	// The number of results per page. Maximum value is 1000.
	PerPage param.Field[int64] `query:"per_page"`
	// Search term for filtering records. Behavior may change.
	Search param.Field[string] `query:"search"`
}

// URLQuery serializes [SettingBlockSenderListParams]'s query parameters as
// `url.Values`.
func (r SettingBlockSenderListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// The sorting direction.
type SettingBlockSenderListParamsDirection string

const (
	SettingBlockSenderListParamsDirectionAsc  SettingBlockSenderListParamsDirection = "asc"
	SettingBlockSenderListParamsDirectionDesc SettingBlockSenderListParamsDirection = "desc"
)

func (r SettingBlockSenderListParamsDirection) IsKnown() bool {
	switch r {
	case SettingBlockSenderListParamsDirectionAsc, SettingBlockSenderListParamsDirectionDesc:
		return true
	}
	return false
}

// Field to sort by.
type SettingBlockSenderListParamsOrder string

const (
	SettingBlockSenderListParamsOrderPattern   SettingBlockSenderListParamsOrder = "pattern"
	SettingBlockSenderListParamsOrderCreatedAt SettingBlockSenderListParamsOrder = "created_at"
)

func (r SettingBlockSenderListParamsOrder) IsKnown() bool {
	switch r {
	case SettingBlockSenderListParamsOrderPattern, SettingBlockSenderListParamsOrderCreatedAt:
		return true
	}
	return false
}

// Filter by pattern type.
type SettingBlockSenderListParamsPatternType string

const (
	SettingBlockSenderListParamsPatternTypeEmail   SettingBlockSenderListParamsPatternType = "EMAIL"
	SettingBlockSenderListParamsPatternTypeDomain  SettingBlockSenderListParamsPatternType = "DOMAIN"
	SettingBlockSenderListParamsPatternTypeIP      SettingBlockSenderListParamsPatternType = "IP"
	SettingBlockSenderListParamsPatternTypeUnknown SettingBlockSenderListParamsPatternType = "UNKNOWN"
)

func (r SettingBlockSenderListParamsPatternType) IsKnown() bool {
	switch r {
	case SettingBlockSenderListParamsPatternTypeEmail, SettingBlockSenderListParamsPatternTypeDomain, SettingBlockSenderListParamsPatternTypeIP, SettingBlockSenderListParamsPatternTypeUnknown:
		return true
	}
	return false
}

type SettingBlockSenderDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type SettingBlockSenderDeleteResponseEnvelope struct {
	Errors   []SettingBlockSenderDeleteResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []SettingBlockSenderDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success SettingBlockSenderDeleteResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  SettingBlockSenderDeleteResponse                `json:"result"`
	JSON    settingBlockSenderDeleteResponseEnvelopeJSON    `json:"-"`
}

// settingBlockSenderDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingBlockSenderDeleteResponseEnvelope]
type settingBlockSenderDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingBlockSenderDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBlockSenderDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingBlockSenderDeleteResponseEnvelopeErrors struct {
	Code             int64                                                `json:"code" api:"required"`
	Message          string                                               `json:"message" api:"required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           SettingBlockSenderDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             settingBlockSenderDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// settingBlockSenderDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SettingBlockSenderDeleteResponseEnvelopeErrors]
type settingBlockSenderDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingBlockSenderDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBlockSenderDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingBlockSenderDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    settingBlockSenderDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// settingBlockSenderDeleteResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [SettingBlockSenderDeleteResponseEnvelopeErrorsSource]
type settingBlockSenderDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingBlockSenderDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBlockSenderDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SettingBlockSenderDeleteResponseEnvelopeMessages struct {
	Code             int64                                                  `json:"code" api:"required"`
	Message          string                                                 `json:"message" api:"required"`
	DocumentationURL string                                                 `json:"documentation_url"`
	Source           SettingBlockSenderDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             settingBlockSenderDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// settingBlockSenderDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingBlockSenderDeleteResponseEnvelopeMessages]
type settingBlockSenderDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingBlockSenderDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBlockSenderDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingBlockSenderDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                     `json:"pointer"`
	JSON    settingBlockSenderDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// settingBlockSenderDeleteResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [SettingBlockSenderDeleteResponseEnvelopeMessagesSource]
type settingBlockSenderDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingBlockSenderDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBlockSenderDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SettingBlockSenderDeleteResponseEnvelopeSuccess bool

const (
	SettingBlockSenderDeleteResponseEnvelopeSuccessTrue SettingBlockSenderDeleteResponseEnvelopeSuccess = true
)

func (r SettingBlockSenderDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingBlockSenderDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SettingBlockSenderEditParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	Comments  param.Field[string] `json:"comments"`
	IsRegex   param.Field[bool]   `json:"is_regex"`
	Pattern   param.Field[string] `json:"pattern"`
	// Type of pattern matching. Note: UNKNOWN is deprecated and cannot be used when
	// creating or updating policies, but may be returned for existing entries.
	PatternType param.Field[SettingBlockSenderEditParamsPatternType] `json:"pattern_type"`
}

func (r SettingBlockSenderEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of pattern matching. Note: UNKNOWN is deprecated and cannot be used when
// creating or updating policies, but may be returned for existing entries.
type SettingBlockSenderEditParamsPatternType string

const (
	SettingBlockSenderEditParamsPatternTypeEmail   SettingBlockSenderEditParamsPatternType = "EMAIL"
	SettingBlockSenderEditParamsPatternTypeDomain  SettingBlockSenderEditParamsPatternType = "DOMAIN"
	SettingBlockSenderEditParamsPatternTypeIP      SettingBlockSenderEditParamsPatternType = "IP"
	SettingBlockSenderEditParamsPatternTypeUnknown SettingBlockSenderEditParamsPatternType = "UNKNOWN"
)

func (r SettingBlockSenderEditParamsPatternType) IsKnown() bool {
	switch r {
	case SettingBlockSenderEditParamsPatternTypeEmail, SettingBlockSenderEditParamsPatternTypeDomain, SettingBlockSenderEditParamsPatternTypeIP, SettingBlockSenderEditParamsPatternTypeUnknown:
		return true
	}
	return false
}

type SettingBlockSenderEditResponseEnvelope struct {
	Errors   []SettingBlockSenderEditResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []SettingBlockSenderEditResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success SettingBlockSenderEditResponseEnvelopeSuccess `json:"success" api:"required"`
	// A blocked sender pattern
	Result SettingBlockSenderEditResponse             `json:"result"`
	JSON   settingBlockSenderEditResponseEnvelopeJSON `json:"-"`
}

// settingBlockSenderEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingBlockSenderEditResponseEnvelope]
type settingBlockSenderEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingBlockSenderEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBlockSenderEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingBlockSenderEditResponseEnvelopeErrors struct {
	Code             int64                                              `json:"code" api:"required"`
	Message          string                                             `json:"message" api:"required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           SettingBlockSenderEditResponseEnvelopeErrorsSource `json:"source"`
	JSON             settingBlockSenderEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// settingBlockSenderEditResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SettingBlockSenderEditResponseEnvelopeErrors]
type settingBlockSenderEditResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingBlockSenderEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBlockSenderEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingBlockSenderEditResponseEnvelopeErrorsSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    settingBlockSenderEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// settingBlockSenderEditResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [SettingBlockSenderEditResponseEnvelopeErrorsSource]
type settingBlockSenderEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingBlockSenderEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBlockSenderEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SettingBlockSenderEditResponseEnvelopeMessages struct {
	Code             int64                                                `json:"code" api:"required"`
	Message          string                                               `json:"message" api:"required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           SettingBlockSenderEditResponseEnvelopeMessagesSource `json:"source"`
	JSON             settingBlockSenderEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// settingBlockSenderEditResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [SettingBlockSenderEditResponseEnvelopeMessages]
type settingBlockSenderEditResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingBlockSenderEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBlockSenderEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingBlockSenderEditResponseEnvelopeMessagesSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    settingBlockSenderEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// settingBlockSenderEditResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [SettingBlockSenderEditResponseEnvelopeMessagesSource]
type settingBlockSenderEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingBlockSenderEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBlockSenderEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SettingBlockSenderEditResponseEnvelopeSuccess bool

const (
	SettingBlockSenderEditResponseEnvelopeSuccessTrue SettingBlockSenderEditResponseEnvelopeSuccess = true
)

func (r SettingBlockSenderEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingBlockSenderEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SettingBlockSenderGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type SettingBlockSenderGetResponseEnvelope struct {
	Errors   []SettingBlockSenderGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []SettingBlockSenderGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success SettingBlockSenderGetResponseEnvelopeSuccess `json:"success" api:"required"`
	// A blocked sender pattern
	Result SettingBlockSenderGetResponse             `json:"result"`
	JSON   settingBlockSenderGetResponseEnvelopeJSON `json:"-"`
}

// settingBlockSenderGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingBlockSenderGetResponseEnvelope]
type settingBlockSenderGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingBlockSenderGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBlockSenderGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingBlockSenderGetResponseEnvelopeErrors struct {
	Code             int64                                             `json:"code" api:"required"`
	Message          string                                            `json:"message" api:"required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           SettingBlockSenderGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             settingBlockSenderGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// settingBlockSenderGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [SettingBlockSenderGetResponseEnvelopeErrors]
type settingBlockSenderGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingBlockSenderGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBlockSenderGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingBlockSenderGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    settingBlockSenderGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// settingBlockSenderGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [SettingBlockSenderGetResponseEnvelopeErrorsSource]
type settingBlockSenderGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingBlockSenderGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBlockSenderGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SettingBlockSenderGetResponseEnvelopeMessages struct {
	Code             int64                                               `json:"code" api:"required"`
	Message          string                                              `json:"message" api:"required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           SettingBlockSenderGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             settingBlockSenderGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// settingBlockSenderGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SettingBlockSenderGetResponseEnvelopeMessages]
type settingBlockSenderGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingBlockSenderGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBlockSenderGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingBlockSenderGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    settingBlockSenderGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// settingBlockSenderGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [SettingBlockSenderGetResponseEnvelopeMessagesSource]
type settingBlockSenderGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingBlockSenderGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingBlockSenderGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SettingBlockSenderGetResponseEnvelopeSuccess bool

const (
	SettingBlockSenderGetResponseEnvelopeSuccessTrue SettingBlockSenderGetResponseEnvelopeSuccess = true
)

func (r SettingBlockSenderGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingBlockSenderGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
