// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
)

// ThreatSignalSkillService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreatSignalSkillService] method instead.
type ThreatSignalSkillService struct {
	Options       []option.RequestOption
	TagCategories *ThreatSignalSkillTagCategoryService
}

// NewThreatSignalSkillService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewThreatSignalSkillService(opts ...option.RequestOption) (r *ThreatSignalSkillService) {
	r = &ThreatSignalSkillService{}
	r.Options = opts
	r.TagCategories = NewThreatSignalSkillTagCategoryService(opts...)
	return
}

// Creates a custom AI skill for the account.
func (r *ThreatSignalSkillService) New(ctx context.Context, params ThreatSignalSkillNewParams, opts ...option.RequestOption) (res *ThreatSignalSkillNewResponse, err error) {
	var env ThreatSignalSkillNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/v2/threat-signals/skills", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Lists the default and custom skills available to the account.
func (r *ThreatSignalSkillService) List(ctx context.Context, params ThreatSignalSkillListParams, opts ...option.RequestOption) (res *pagination.V4PagePagination[ThreatSignalSkillListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/v2/threat-signals/skills", params.AccountID)
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

// Lists the default and custom skills available to the account.
func (r *ThreatSignalSkillService) ListAutoPaging(ctx context.Context, params ThreatSignalSkillListParams, opts ...option.RequestOption) *pagination.V4PagePaginationAutoPager[ThreatSignalSkillListResponse] {
	return pagination.NewV4PagePaginationAutoPager(r.List(ctx, params, opts...))
}

// Deletes a custom skill. Default skills cannot be deleted.
func (r *ThreatSignalSkillService) Delete(ctx context.Context, skillID string, body ThreatSignalSkillDeleteParams, opts ...option.RequestOption) (res *ThreatSignalSkillDeleteResponse, err error) {
	var env ThreatSignalSkillDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if skillID == "" {
		err = errors.New("missing required skill_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/v2/threat-signals/skills/%s", body.AccountID, skillID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Updates a custom skill. Default skills are read-only.
func (r *ThreatSignalSkillService) Edit(ctx context.Context, skillID string, params ThreatSignalSkillEditParams, opts ...option.RequestOption) (res *ThreatSignalSkillEditResponse, err error) {
	var env ThreatSignalSkillEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if skillID == "" {
		err = errors.New("missing required skill_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/v2/threat-signals/skills/%s", params.AccountID, skillID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieves a default or custom skill by ID.
func (r *ThreatSignalSkillService) Get(ctx context.Context, skillID string, query ThreatSignalSkillGetParams, opts ...option.RequestOption) (res *ThreatSignalSkillGetResponse, err error) {
	var env ThreatSignalSkillGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if skillID == "" {
		err = errors.New("missing required skill_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/v2/threat-signals/skills/%s", query.AccountID, skillID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type ThreatSignalSkillNewResponse struct {
	ID string `json:"id" api:"required"`
	// JSON-encoded skill configuration. Always null for default skills.
	Config    string `json:"config" api:"required,nullable"`
	CreatedAt string `json:"created_at" api:"required"`
	// 1 when active, 0 when inactive.
	IsActive int64  `json:"is_active" api:"required"`
	Name     string `json:"name" api:"required"`
	// JSON-encoded JSON Schema the skill output must satisfy.
	OutputSchema string `json:"output_schema" api:"required,nullable"`
	Prompt       string `json:"prompt" api:"required"`
	// `default` for Cloudforce One managed skills (read-only), `custom` for account
	// skills.
	Source    ThreatSignalSkillNewResponseSource `json:"source" api:"required"`
	Type      string                             `json:"type" api:"required"`
	UpdatedAt string                             `json:"updated_at" api:"required"`
	JSON      threatSignalSkillNewResponseJSON   `json:"-"`
}

// threatSignalSkillNewResponseJSON contains the JSON metadata for the struct
// [ThreatSignalSkillNewResponse]
type threatSignalSkillNewResponseJSON struct {
	ID           apijson.Field
	Config       apijson.Field
	CreatedAt    apijson.Field
	IsActive     apijson.Field
	Name         apijson.Field
	OutputSchema apijson.Field
	Prompt       apijson.Field
	Source       apijson.Field
	Type         apijson.Field
	UpdatedAt    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ThreatSignalSkillNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalSkillNewResponseJSON) RawJSON() string {
	return r.raw
}

// `default` for Cloudforce One managed skills (read-only), `custom` for account
// skills.
type ThreatSignalSkillNewResponseSource string

const (
	ThreatSignalSkillNewResponseSourceDefault ThreatSignalSkillNewResponseSource = "default"
	ThreatSignalSkillNewResponseSourceCustom  ThreatSignalSkillNewResponseSource = "custom"
)

func (r ThreatSignalSkillNewResponseSource) IsKnown() bool {
	switch r {
	case ThreatSignalSkillNewResponseSourceDefault, ThreatSignalSkillNewResponseSourceCustom:
		return true
	}
	return false
}

type ThreatSignalSkillListResponse struct {
	// Number of skills on this page.
	Count int64 `json:"count" api:"required"`
	// Whether the authenticated account may access custom-skill capabilities under
	// Stakeout's Threat Signals access-mode policy. This is a policy availability
	// indicator, not a row-existence indicator. False for threat_signals_only mode;
	// true for entitled, allowlisted, cfone_internal, and service modes.
	CustomSkillsAvailable bool                                 `json:"custom_skills_available" api:"required"`
	Page                  int64                                `json:"page" api:"required"`
	PerPage               int64                                `json:"per_page" api:"required"`
	Skills                []ThreatSignalSkillListResponseSkill `json:"skills" api:"required"`
	TotalCount            int64                                `json:"total_count" api:"required"`
	JSON                  threatSignalSkillListResponseJSON    `json:"-"`
}

// threatSignalSkillListResponseJSON contains the JSON metadata for the struct
// [ThreatSignalSkillListResponse]
type threatSignalSkillListResponseJSON struct {
	Count                 apijson.Field
	CustomSkillsAvailable apijson.Field
	Page                  apijson.Field
	PerPage               apijson.Field
	Skills                apijson.Field
	TotalCount            apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ThreatSignalSkillListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalSkillListResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalSkillListResponseSkill struct {
	ID string `json:"id" api:"required"`
	// JSON-encoded skill configuration. Always null for default skills.
	Config    string `json:"config" api:"required,nullable"`
	CreatedAt string `json:"created_at" api:"required"`
	// 1 when active, 0 when inactive.
	IsActive int64  `json:"is_active" api:"required"`
	Name     string `json:"name" api:"required"`
	// JSON-encoded JSON Schema the skill output must satisfy.
	OutputSchema string `json:"output_schema" api:"required,nullable"`
	Prompt       string `json:"prompt" api:"required"`
	// `default` for Cloudforce One managed skills (read-only), `custom` for account
	// skills.
	Source    ThreatSignalSkillListResponseSkillsSource `json:"source" api:"required"`
	Type      string                                    `json:"type" api:"required"`
	UpdatedAt string                                    `json:"updated_at" api:"required"`
	JSON      threatSignalSkillListResponseSkillJSON    `json:"-"`
}

// threatSignalSkillListResponseSkillJSON contains the JSON metadata for the struct
// [ThreatSignalSkillListResponseSkill]
type threatSignalSkillListResponseSkillJSON struct {
	ID           apijson.Field
	Config       apijson.Field
	CreatedAt    apijson.Field
	IsActive     apijson.Field
	Name         apijson.Field
	OutputSchema apijson.Field
	Prompt       apijson.Field
	Source       apijson.Field
	Type         apijson.Field
	UpdatedAt    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ThreatSignalSkillListResponseSkill) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalSkillListResponseSkillJSON) RawJSON() string {
	return r.raw
}

// `default` for Cloudforce One managed skills (read-only), `custom` for account
// skills.
type ThreatSignalSkillListResponseSkillsSource string

const (
	ThreatSignalSkillListResponseSkillsSourceDefault ThreatSignalSkillListResponseSkillsSource = "default"
	ThreatSignalSkillListResponseSkillsSourceCustom  ThreatSignalSkillListResponseSkillsSource = "custom"
)

func (r ThreatSignalSkillListResponseSkillsSource) IsKnown() bool {
	switch r {
	case ThreatSignalSkillListResponseSkillsSourceDefault, ThreatSignalSkillListResponseSkillsSourceCustom:
		return true
	}
	return false
}

type ThreatSignalSkillDeleteResponse struct {
	ID string `json:"id" api:"required"`
	// JSON-encoded skill configuration. Always null for default skills.
	Config    string `json:"config" api:"required,nullable"`
	CreatedAt string `json:"created_at" api:"required"`
	// 1 when active, 0 when inactive.
	IsActive int64  `json:"is_active" api:"required"`
	Name     string `json:"name" api:"required"`
	// JSON-encoded JSON Schema the skill output must satisfy.
	OutputSchema string `json:"output_schema" api:"required,nullable"`
	Prompt       string `json:"prompt" api:"required"`
	// `default` for Cloudforce One managed skills (read-only), `custom` for account
	// skills.
	Source    ThreatSignalSkillDeleteResponseSource `json:"source" api:"required"`
	Type      string                                `json:"type" api:"required"`
	UpdatedAt string                                `json:"updated_at" api:"required"`
	JSON      threatSignalSkillDeleteResponseJSON   `json:"-"`
}

// threatSignalSkillDeleteResponseJSON contains the JSON metadata for the struct
// [ThreatSignalSkillDeleteResponse]
type threatSignalSkillDeleteResponseJSON struct {
	ID           apijson.Field
	Config       apijson.Field
	CreatedAt    apijson.Field
	IsActive     apijson.Field
	Name         apijson.Field
	OutputSchema apijson.Field
	Prompt       apijson.Field
	Source       apijson.Field
	Type         apijson.Field
	UpdatedAt    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ThreatSignalSkillDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalSkillDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// `default` for Cloudforce One managed skills (read-only), `custom` for account
// skills.
type ThreatSignalSkillDeleteResponseSource string

const (
	ThreatSignalSkillDeleteResponseSourceDefault ThreatSignalSkillDeleteResponseSource = "default"
	ThreatSignalSkillDeleteResponseSourceCustom  ThreatSignalSkillDeleteResponseSource = "custom"
)

func (r ThreatSignalSkillDeleteResponseSource) IsKnown() bool {
	switch r {
	case ThreatSignalSkillDeleteResponseSourceDefault, ThreatSignalSkillDeleteResponseSourceCustom:
		return true
	}
	return false
}

type ThreatSignalSkillEditResponse struct {
	ID string `json:"id" api:"required"`
	// JSON-encoded skill configuration. Always null for default skills.
	Config    string `json:"config" api:"required,nullable"`
	CreatedAt string `json:"created_at" api:"required"`
	// 1 when active, 0 when inactive.
	IsActive int64  `json:"is_active" api:"required"`
	Name     string `json:"name" api:"required"`
	// JSON-encoded JSON Schema the skill output must satisfy.
	OutputSchema string `json:"output_schema" api:"required,nullable"`
	Prompt       string `json:"prompt" api:"required"`
	// `default` for Cloudforce One managed skills (read-only), `custom` for account
	// skills.
	Source    ThreatSignalSkillEditResponseSource `json:"source" api:"required"`
	Type      string                              `json:"type" api:"required"`
	UpdatedAt string                              `json:"updated_at" api:"required"`
	JSON      threatSignalSkillEditResponseJSON   `json:"-"`
}

// threatSignalSkillEditResponseJSON contains the JSON metadata for the struct
// [ThreatSignalSkillEditResponse]
type threatSignalSkillEditResponseJSON struct {
	ID           apijson.Field
	Config       apijson.Field
	CreatedAt    apijson.Field
	IsActive     apijson.Field
	Name         apijson.Field
	OutputSchema apijson.Field
	Prompt       apijson.Field
	Source       apijson.Field
	Type         apijson.Field
	UpdatedAt    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ThreatSignalSkillEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalSkillEditResponseJSON) RawJSON() string {
	return r.raw
}

// `default` for Cloudforce One managed skills (read-only), `custom` for account
// skills.
type ThreatSignalSkillEditResponseSource string

const (
	ThreatSignalSkillEditResponseSourceDefault ThreatSignalSkillEditResponseSource = "default"
	ThreatSignalSkillEditResponseSourceCustom  ThreatSignalSkillEditResponseSource = "custom"
)

func (r ThreatSignalSkillEditResponseSource) IsKnown() bool {
	switch r {
	case ThreatSignalSkillEditResponseSourceDefault, ThreatSignalSkillEditResponseSourceCustom:
		return true
	}
	return false
}

type ThreatSignalSkillGetResponse struct {
	ID string `json:"id" api:"required"`
	// JSON-encoded skill configuration. Always null for default skills.
	Config    string `json:"config" api:"required,nullable"`
	CreatedAt string `json:"created_at" api:"required"`
	// 1 when active, 0 when inactive.
	IsActive int64  `json:"is_active" api:"required"`
	Name     string `json:"name" api:"required"`
	// JSON-encoded JSON Schema the skill output must satisfy.
	OutputSchema string `json:"output_schema" api:"required,nullable"`
	Prompt       string `json:"prompt" api:"required"`
	// `default` for Cloudforce One managed skills (read-only), `custom` for account
	// skills.
	Source    ThreatSignalSkillGetResponseSource `json:"source" api:"required"`
	Type      string                             `json:"type" api:"required"`
	UpdatedAt string                             `json:"updated_at" api:"required"`
	JSON      threatSignalSkillGetResponseJSON   `json:"-"`
}

// threatSignalSkillGetResponseJSON contains the JSON metadata for the struct
// [ThreatSignalSkillGetResponse]
type threatSignalSkillGetResponseJSON struct {
	ID           apijson.Field
	Config       apijson.Field
	CreatedAt    apijson.Field
	IsActive     apijson.Field
	Name         apijson.Field
	OutputSchema apijson.Field
	Prompt       apijson.Field
	Source       apijson.Field
	Type         apijson.Field
	UpdatedAt    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ThreatSignalSkillGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalSkillGetResponseJSON) RawJSON() string {
	return r.raw
}

// `default` for Cloudforce One managed skills (read-only), `custom` for account
// skills.
type ThreatSignalSkillGetResponseSource string

const (
	ThreatSignalSkillGetResponseSourceDefault ThreatSignalSkillGetResponseSource = "default"
	ThreatSignalSkillGetResponseSourceCustom  ThreatSignalSkillGetResponseSource = "custom"
)

func (r ThreatSignalSkillGetResponseSource) IsKnown() bool {
	switch r {
	case ThreatSignalSkillGetResponseSourceDefault, ThreatSignalSkillGetResponseSourceCustom:
		return true
	}
	return false
}

type ThreatSignalSkillNewParams struct {
	AccountID    param.Field[string]                         `path:"account_id" api:"required"`
	Name         param.Field[string]                         `json:"name" api:"required"`
	OutputSchema param.Field[string]                         `json:"output_schema" api:"required"`
	Prompt       param.Field[string]                         `json:"prompt" api:"required"`
	Type         param.Field[ThreatSignalSkillNewParamsType] `json:"type" api:"required"`
}

func (r ThreatSignalSkillNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ThreatSignalSkillNewParamsType string

const (
	ThreatSignalSkillNewParamsTypeSummary ThreatSignalSkillNewParamsType = "summary"
	ThreatSignalSkillNewParamsTypeTags    ThreatSignalSkillNewParamsType = "tags"
)

func (r ThreatSignalSkillNewParamsType) IsKnown() bool {
	switch r {
	case ThreatSignalSkillNewParamsTypeSummary, ThreatSignalSkillNewParamsTypeTags:
		return true
	}
	return false
}

type ThreatSignalSkillNewResponseEnvelope struct {
	Errors   []ThreatSignalSkillNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ThreatSignalSkillNewResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   ThreatSignalSkillNewResponse                   `json:"result" api:"required"`
	Success  ThreatSignalSkillNewResponseEnvelopeSuccess    `json:"success" api:"required"`
	JSON     threatSignalSkillNewResponseEnvelopeJSON       `json:"-"`
}

// threatSignalSkillNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [ThreatSignalSkillNewResponseEnvelope]
type threatSignalSkillNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalSkillNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalSkillNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalSkillNewResponseEnvelopeErrors struct {
	Message string                                         `json:"message" api:"required"`
	JSON    threatSignalSkillNewResponseEnvelopeErrorsJSON `json:"-"`
}

// threatSignalSkillNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ThreatSignalSkillNewResponseEnvelopeErrors]
type threatSignalSkillNewResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalSkillNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalSkillNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalSkillNewResponseEnvelopeMessages struct {
	Message  string                                             `json:"message" api:"required"`
	Code     float64                                            `json:"code"`
	Expected string                                             `json:"expected"`
	Path     []string                                           `json:"path"`
	Reason   ThreatSignalSkillNewResponseEnvelopeMessagesReason `json:"reason"`
	JSON     threatSignalSkillNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// threatSignalSkillNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ThreatSignalSkillNewResponseEnvelopeMessages]
type threatSignalSkillNewResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Expected    apijson.Field
	Path        apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalSkillNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalSkillNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalSkillNewResponseEnvelopeMessagesReason string

const (
	ThreatSignalSkillNewResponseEnvelopeMessagesReasonFreeCustomFeedLimit              ThreatSignalSkillNewResponseEnvelopeMessagesReason = "free_custom_feed_limit"
	ThreatSignalSkillNewResponseEnvelopeMessagesReasonFreeCustomSkillsDisabled         ThreatSignalSkillNewResponseEnvelopeMessagesReason = "free_custom_skills_disabled"
	ThreatSignalSkillNewResponseEnvelopeMessagesReasonFreeTierReconciliationInProgress ThreatSignalSkillNewResponseEnvelopeMessagesReason = "free_tier_reconciliation_in_progress"
	ThreatSignalSkillNewResponseEnvelopeMessagesReasonFreeTierReconciliationFailed     ThreatSignalSkillNewResponseEnvelopeMessagesReason = "free_tier_reconciliation_failed"
	ThreatSignalSkillNewResponseEnvelopeMessagesReasonRawContentExpired                ThreatSignalSkillNewResponseEnvelopeMessagesReason = "raw_content_expired"
	ThreatSignalSkillNewResponseEnvelopeMessagesReasonCuratedVisibilityChanged         ThreatSignalSkillNewResponseEnvelopeMessagesReason = "curated_visibility_changed"
)

func (r ThreatSignalSkillNewResponseEnvelopeMessagesReason) IsKnown() bool {
	switch r {
	case ThreatSignalSkillNewResponseEnvelopeMessagesReasonFreeCustomFeedLimit, ThreatSignalSkillNewResponseEnvelopeMessagesReasonFreeCustomSkillsDisabled, ThreatSignalSkillNewResponseEnvelopeMessagesReasonFreeTierReconciliationInProgress, ThreatSignalSkillNewResponseEnvelopeMessagesReasonFreeTierReconciliationFailed, ThreatSignalSkillNewResponseEnvelopeMessagesReasonRawContentExpired, ThreatSignalSkillNewResponseEnvelopeMessagesReasonCuratedVisibilityChanged:
		return true
	}
	return false
}

type ThreatSignalSkillNewResponseEnvelopeSuccess bool

const (
	ThreatSignalSkillNewResponseEnvelopeSuccessTrue ThreatSignalSkillNewResponseEnvelopeSuccess = true
)

func (r ThreatSignalSkillNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ThreatSignalSkillNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ThreatSignalSkillListParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	Page      param.Field[int64]  `query:"page"`
	PerPage   param.Field[int64]  `query:"per_page"`
}

// URLQuery serializes [ThreatSignalSkillListParams]'s query parameters as
// `url.Values`.
func (r ThreatSignalSkillListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ThreatSignalSkillDeleteParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type ThreatSignalSkillDeleteResponseEnvelope struct {
	Errors   []ThreatSignalSkillDeleteResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ThreatSignalSkillDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   ThreatSignalSkillDeleteResponse                   `json:"result" api:"required"`
	Success  ThreatSignalSkillDeleteResponseEnvelopeSuccess    `json:"success" api:"required"`
	JSON     threatSignalSkillDeleteResponseEnvelopeJSON       `json:"-"`
}

// threatSignalSkillDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [ThreatSignalSkillDeleteResponseEnvelope]
type threatSignalSkillDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalSkillDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalSkillDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalSkillDeleteResponseEnvelopeErrors struct {
	Message string                                            `json:"message" api:"required"`
	JSON    threatSignalSkillDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// threatSignalSkillDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ThreatSignalSkillDeleteResponseEnvelopeErrors]
type threatSignalSkillDeleteResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalSkillDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalSkillDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalSkillDeleteResponseEnvelopeMessages struct {
	Message  string                                                `json:"message" api:"required"`
	Code     float64                                               `json:"code"`
	Expected string                                                `json:"expected"`
	Path     []string                                              `json:"path"`
	Reason   ThreatSignalSkillDeleteResponseEnvelopeMessagesReason `json:"reason"`
	JSON     threatSignalSkillDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// threatSignalSkillDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ThreatSignalSkillDeleteResponseEnvelopeMessages]
type threatSignalSkillDeleteResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Expected    apijson.Field
	Path        apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalSkillDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalSkillDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalSkillDeleteResponseEnvelopeMessagesReason string

const (
	ThreatSignalSkillDeleteResponseEnvelopeMessagesReasonFreeCustomFeedLimit              ThreatSignalSkillDeleteResponseEnvelopeMessagesReason = "free_custom_feed_limit"
	ThreatSignalSkillDeleteResponseEnvelopeMessagesReasonFreeCustomSkillsDisabled         ThreatSignalSkillDeleteResponseEnvelopeMessagesReason = "free_custom_skills_disabled"
	ThreatSignalSkillDeleteResponseEnvelopeMessagesReasonFreeTierReconciliationInProgress ThreatSignalSkillDeleteResponseEnvelopeMessagesReason = "free_tier_reconciliation_in_progress"
	ThreatSignalSkillDeleteResponseEnvelopeMessagesReasonFreeTierReconciliationFailed     ThreatSignalSkillDeleteResponseEnvelopeMessagesReason = "free_tier_reconciliation_failed"
	ThreatSignalSkillDeleteResponseEnvelopeMessagesReasonRawContentExpired                ThreatSignalSkillDeleteResponseEnvelopeMessagesReason = "raw_content_expired"
	ThreatSignalSkillDeleteResponseEnvelopeMessagesReasonCuratedVisibilityChanged         ThreatSignalSkillDeleteResponseEnvelopeMessagesReason = "curated_visibility_changed"
)

func (r ThreatSignalSkillDeleteResponseEnvelopeMessagesReason) IsKnown() bool {
	switch r {
	case ThreatSignalSkillDeleteResponseEnvelopeMessagesReasonFreeCustomFeedLimit, ThreatSignalSkillDeleteResponseEnvelopeMessagesReasonFreeCustomSkillsDisabled, ThreatSignalSkillDeleteResponseEnvelopeMessagesReasonFreeTierReconciliationInProgress, ThreatSignalSkillDeleteResponseEnvelopeMessagesReasonFreeTierReconciliationFailed, ThreatSignalSkillDeleteResponseEnvelopeMessagesReasonRawContentExpired, ThreatSignalSkillDeleteResponseEnvelopeMessagesReasonCuratedVisibilityChanged:
		return true
	}
	return false
}

type ThreatSignalSkillDeleteResponseEnvelopeSuccess bool

const (
	ThreatSignalSkillDeleteResponseEnvelopeSuccessTrue ThreatSignalSkillDeleteResponseEnvelopeSuccess = true
)

func (r ThreatSignalSkillDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ThreatSignalSkillDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ThreatSignalSkillEditParams struct {
	AccountID    param.Field[string] `path:"account_id" api:"required"`
	Config       param.Field[string] `json:"config"`
	IsActive     param.Field[bool]   `json:"is_active"`
	Name         param.Field[string] `json:"name"`
	OutputSchema param.Field[string] `json:"output_schema"`
	Prompt       param.Field[string] `json:"prompt"`
}

func (r ThreatSignalSkillEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ThreatSignalSkillEditResponseEnvelope struct {
	Errors   []ThreatSignalSkillEditResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ThreatSignalSkillEditResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   ThreatSignalSkillEditResponse                   `json:"result" api:"required"`
	Success  ThreatSignalSkillEditResponseEnvelopeSuccess    `json:"success" api:"required"`
	JSON     threatSignalSkillEditResponseEnvelopeJSON       `json:"-"`
}

// threatSignalSkillEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [ThreatSignalSkillEditResponseEnvelope]
type threatSignalSkillEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalSkillEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalSkillEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalSkillEditResponseEnvelopeErrors struct {
	Message string                                          `json:"message" api:"required"`
	JSON    threatSignalSkillEditResponseEnvelopeErrorsJSON `json:"-"`
}

// threatSignalSkillEditResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ThreatSignalSkillEditResponseEnvelopeErrors]
type threatSignalSkillEditResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalSkillEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalSkillEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalSkillEditResponseEnvelopeMessages struct {
	Message  string                                              `json:"message" api:"required"`
	Code     float64                                             `json:"code"`
	Expected string                                              `json:"expected"`
	Path     []string                                            `json:"path"`
	Reason   ThreatSignalSkillEditResponseEnvelopeMessagesReason `json:"reason"`
	JSON     threatSignalSkillEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// threatSignalSkillEditResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ThreatSignalSkillEditResponseEnvelopeMessages]
type threatSignalSkillEditResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Expected    apijson.Field
	Path        apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalSkillEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalSkillEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalSkillEditResponseEnvelopeMessagesReason string

const (
	ThreatSignalSkillEditResponseEnvelopeMessagesReasonFreeCustomFeedLimit              ThreatSignalSkillEditResponseEnvelopeMessagesReason = "free_custom_feed_limit"
	ThreatSignalSkillEditResponseEnvelopeMessagesReasonFreeCustomSkillsDisabled         ThreatSignalSkillEditResponseEnvelopeMessagesReason = "free_custom_skills_disabled"
	ThreatSignalSkillEditResponseEnvelopeMessagesReasonFreeTierReconciliationInProgress ThreatSignalSkillEditResponseEnvelopeMessagesReason = "free_tier_reconciliation_in_progress"
	ThreatSignalSkillEditResponseEnvelopeMessagesReasonFreeTierReconciliationFailed     ThreatSignalSkillEditResponseEnvelopeMessagesReason = "free_tier_reconciliation_failed"
	ThreatSignalSkillEditResponseEnvelopeMessagesReasonRawContentExpired                ThreatSignalSkillEditResponseEnvelopeMessagesReason = "raw_content_expired"
	ThreatSignalSkillEditResponseEnvelopeMessagesReasonCuratedVisibilityChanged         ThreatSignalSkillEditResponseEnvelopeMessagesReason = "curated_visibility_changed"
)

func (r ThreatSignalSkillEditResponseEnvelopeMessagesReason) IsKnown() bool {
	switch r {
	case ThreatSignalSkillEditResponseEnvelopeMessagesReasonFreeCustomFeedLimit, ThreatSignalSkillEditResponseEnvelopeMessagesReasonFreeCustomSkillsDisabled, ThreatSignalSkillEditResponseEnvelopeMessagesReasonFreeTierReconciliationInProgress, ThreatSignalSkillEditResponseEnvelopeMessagesReasonFreeTierReconciliationFailed, ThreatSignalSkillEditResponseEnvelopeMessagesReasonRawContentExpired, ThreatSignalSkillEditResponseEnvelopeMessagesReasonCuratedVisibilityChanged:
		return true
	}
	return false
}

type ThreatSignalSkillEditResponseEnvelopeSuccess bool

const (
	ThreatSignalSkillEditResponseEnvelopeSuccessTrue ThreatSignalSkillEditResponseEnvelopeSuccess = true
)

func (r ThreatSignalSkillEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ThreatSignalSkillEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ThreatSignalSkillGetParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type ThreatSignalSkillGetResponseEnvelope struct {
	Errors   []ThreatSignalSkillGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ThreatSignalSkillGetResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   ThreatSignalSkillGetResponse                   `json:"result" api:"required"`
	Success  ThreatSignalSkillGetResponseEnvelopeSuccess    `json:"success" api:"required"`
	JSON     threatSignalSkillGetResponseEnvelopeJSON       `json:"-"`
}

// threatSignalSkillGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ThreatSignalSkillGetResponseEnvelope]
type threatSignalSkillGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalSkillGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalSkillGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalSkillGetResponseEnvelopeErrors struct {
	Message string                                         `json:"message" api:"required"`
	JSON    threatSignalSkillGetResponseEnvelopeErrorsJSON `json:"-"`
}

// threatSignalSkillGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ThreatSignalSkillGetResponseEnvelopeErrors]
type threatSignalSkillGetResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalSkillGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalSkillGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalSkillGetResponseEnvelopeMessages struct {
	Message  string                                             `json:"message" api:"required"`
	Code     float64                                            `json:"code"`
	Expected string                                             `json:"expected"`
	Path     []string                                           `json:"path"`
	Reason   ThreatSignalSkillGetResponseEnvelopeMessagesReason `json:"reason"`
	JSON     threatSignalSkillGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// threatSignalSkillGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ThreatSignalSkillGetResponseEnvelopeMessages]
type threatSignalSkillGetResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Expected    apijson.Field
	Path        apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalSkillGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalSkillGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalSkillGetResponseEnvelopeMessagesReason string

const (
	ThreatSignalSkillGetResponseEnvelopeMessagesReasonFreeCustomFeedLimit              ThreatSignalSkillGetResponseEnvelopeMessagesReason = "free_custom_feed_limit"
	ThreatSignalSkillGetResponseEnvelopeMessagesReasonFreeCustomSkillsDisabled         ThreatSignalSkillGetResponseEnvelopeMessagesReason = "free_custom_skills_disabled"
	ThreatSignalSkillGetResponseEnvelopeMessagesReasonFreeTierReconciliationInProgress ThreatSignalSkillGetResponseEnvelopeMessagesReason = "free_tier_reconciliation_in_progress"
	ThreatSignalSkillGetResponseEnvelopeMessagesReasonFreeTierReconciliationFailed     ThreatSignalSkillGetResponseEnvelopeMessagesReason = "free_tier_reconciliation_failed"
	ThreatSignalSkillGetResponseEnvelopeMessagesReasonRawContentExpired                ThreatSignalSkillGetResponseEnvelopeMessagesReason = "raw_content_expired"
	ThreatSignalSkillGetResponseEnvelopeMessagesReasonCuratedVisibilityChanged         ThreatSignalSkillGetResponseEnvelopeMessagesReason = "curated_visibility_changed"
)

func (r ThreatSignalSkillGetResponseEnvelopeMessagesReason) IsKnown() bool {
	switch r {
	case ThreatSignalSkillGetResponseEnvelopeMessagesReasonFreeCustomFeedLimit, ThreatSignalSkillGetResponseEnvelopeMessagesReasonFreeCustomSkillsDisabled, ThreatSignalSkillGetResponseEnvelopeMessagesReasonFreeTierReconciliationInProgress, ThreatSignalSkillGetResponseEnvelopeMessagesReasonFreeTierReconciliationFailed, ThreatSignalSkillGetResponseEnvelopeMessagesReasonRawContentExpired, ThreatSignalSkillGetResponseEnvelopeMessagesReasonCuratedVisibilityChanged:
		return true
	}
	return false
}

type ThreatSignalSkillGetResponseEnvelopeSuccess bool

const (
	ThreatSignalSkillGetResponseEnvelopeSuccessTrue ThreatSignalSkillGetResponseEnvelopeSuccess = true
)

func (r ThreatSignalSkillGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ThreatSignalSkillGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
