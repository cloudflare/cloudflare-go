// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// ThreatSignalArticleTagService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreatSignalArticleTagService] method instead.
type ThreatSignalArticleTagService struct {
	Options []option.RequestOption
}

// NewThreatSignalArticleTagService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewThreatSignalArticleTagService(opts ...option.RequestOption) (r *ThreatSignalArticleTagService) {
	r = &ThreatSignalArticleTagService{}
	r.Options = opts
	return
}

// Applies a tag from the account's tag catalog to a Threat Signals article.
func (r *ThreatSignalArticleTagService) New(ctx context.Context, articleID string, params ThreatSignalArticleTagNewParams, opts ...option.RequestOption) (res *ThreatSignalArticleTagNewResponse, err error) {
	var env ThreatSignalArticleTagNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if articleID == "" {
		err = errors.New("missing required article_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/v2/threat-signals/articles/%s/tags", params.AccountID, articleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Removes a tag from a Threat Signals article.
func (r *ThreatSignalArticleTagService) Delete(ctx context.Context, articleID string, tagID string, body ThreatSignalArticleTagDeleteParams, opts ...option.RequestOption) (res *ThreatSignalArticleTagDeleteResponse, err error) {
	var env ThreatSignalArticleTagDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if articleID == "" {
		err = errors.New("missing required article_id parameter")
		return nil, err
	}
	if tagID == "" {
		err = errors.New("missing required tag_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/v2/threat-signals/articles/%s/tags/%s", body.AccountID, articleID, tagID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Runs the default AI tagging skill on an article and replaces its AI-applied
// tags.
func (r *ThreatSignalArticleTagService) Generate(ctx context.Context, articleID string, body ThreatSignalArticleTagGenerateParams, opts ...option.RequestOption) (res *ThreatSignalArticleTagGenerateResponse, err error) {
	var env ThreatSignalArticleTagGenerateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if articleID == "" {
		err = errors.New("missing required article_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/v2/threat-signals/articles/%s/tag", body.AccountID, articleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type ThreatSignalArticleTagNewResponse struct {
	AppliedBy  ThreatSignalArticleTagNewResponseAppliedBy `json:"applied_by" api:"required"`
	CategoryID string                                     `json:"categoryId" api:"required,nullable" format:"uuid"`
	UUID       string                                     `json:"uuid" api:"required" format:"uuid"`
	Value      string                                     `json:"value" api:"required"`
	JSON       threatSignalArticleTagNewResponseJSON      `json:"-"`
}

// threatSignalArticleTagNewResponseJSON contains the JSON metadata for the struct
// [ThreatSignalArticleTagNewResponse]
type threatSignalArticleTagNewResponseJSON struct {
	AppliedBy   apijson.Field
	CategoryID  apijson.Field
	UUID        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalArticleTagNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalArticleTagNewResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalArticleTagNewResponseAppliedBy string

const (
	ThreatSignalArticleTagNewResponseAppliedByAI      ThreatSignalArticleTagNewResponseAppliedBy = "ai"
	ThreatSignalArticleTagNewResponseAppliedByAnalyst ThreatSignalArticleTagNewResponseAppliedBy = "analyst"
	ThreatSignalArticleTagNewResponseAppliedBySystem  ThreatSignalArticleTagNewResponseAppliedBy = "system"
)

func (r ThreatSignalArticleTagNewResponseAppliedBy) IsKnown() bool {
	switch r {
	case ThreatSignalArticleTagNewResponseAppliedByAI, ThreatSignalArticleTagNewResponseAppliedByAnalyst, ThreatSignalArticleTagNewResponseAppliedBySystem:
		return true
	}
	return false
}

type ThreatSignalArticleTagDeleteResponse struct {
	AppliedBy  ThreatSignalArticleTagDeleteResponseAppliedBy `json:"applied_by" api:"required"`
	CategoryID string                                        `json:"categoryId" api:"required,nullable" format:"uuid"`
	UUID       string                                        `json:"uuid" api:"required" format:"uuid"`
	Value      string                                        `json:"value" api:"required"`
	JSON       threatSignalArticleTagDeleteResponseJSON      `json:"-"`
}

// threatSignalArticleTagDeleteResponseJSON contains the JSON metadata for the
// struct [ThreatSignalArticleTagDeleteResponse]
type threatSignalArticleTagDeleteResponseJSON struct {
	AppliedBy   apijson.Field
	CategoryID  apijson.Field
	UUID        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalArticleTagDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalArticleTagDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalArticleTagDeleteResponseAppliedBy string

const (
	ThreatSignalArticleTagDeleteResponseAppliedByAI      ThreatSignalArticleTagDeleteResponseAppliedBy = "ai"
	ThreatSignalArticleTagDeleteResponseAppliedByAnalyst ThreatSignalArticleTagDeleteResponseAppliedBy = "analyst"
	ThreatSignalArticleTagDeleteResponseAppliedBySystem  ThreatSignalArticleTagDeleteResponseAppliedBy = "system"
)

func (r ThreatSignalArticleTagDeleteResponseAppliedBy) IsKnown() bool {
	switch r {
	case ThreatSignalArticleTagDeleteResponseAppliedByAI, ThreatSignalArticleTagDeleteResponseAppliedByAnalyst, ThreatSignalArticleTagDeleteResponseAppliedBySystem:
		return true
	}
	return false
}

type ThreatSignalArticleTagGenerateResponse struct {
	TagSkillVersion string `json:"tag_skill_version" api:"required"`
	// Final hydrated assignment set; may be empty when no applicable tags are
	// selected.
	Tags []ThreatSignalArticleTagGenerateResponseTag `json:"tags" api:"required"`
	JSON threatSignalArticleTagGenerateResponseJSON  `json:"-"`
}

// threatSignalArticleTagGenerateResponseJSON contains the JSON metadata for the
// struct [ThreatSignalArticleTagGenerateResponse]
type threatSignalArticleTagGenerateResponseJSON struct {
	TagSkillVersion apijson.Field
	Tags            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ThreatSignalArticleTagGenerateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalArticleTagGenerateResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalArticleTagGenerateResponseTag struct {
	AppliedBy  ThreatSignalArticleTagGenerateResponseTagsAppliedBy `json:"applied_by" api:"required"`
	CategoryID string                                              `json:"categoryId" api:"required,nullable" format:"uuid"`
	UUID       string                                              `json:"uuid" api:"required" format:"uuid"`
	Value      string                                              `json:"value" api:"required"`
	JSON       threatSignalArticleTagGenerateResponseTagJSON       `json:"-"`
}

// threatSignalArticleTagGenerateResponseTagJSON contains the JSON metadata for the
// struct [ThreatSignalArticleTagGenerateResponseTag]
type threatSignalArticleTagGenerateResponseTagJSON struct {
	AppliedBy   apijson.Field
	CategoryID  apijson.Field
	UUID        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalArticleTagGenerateResponseTag) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalArticleTagGenerateResponseTagJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalArticleTagGenerateResponseTagsAppliedBy string

const (
	ThreatSignalArticleTagGenerateResponseTagsAppliedByAI      ThreatSignalArticleTagGenerateResponseTagsAppliedBy = "ai"
	ThreatSignalArticleTagGenerateResponseTagsAppliedByAnalyst ThreatSignalArticleTagGenerateResponseTagsAppliedBy = "analyst"
	ThreatSignalArticleTagGenerateResponseTagsAppliedBySystem  ThreatSignalArticleTagGenerateResponseTagsAppliedBy = "system"
)

func (r ThreatSignalArticleTagGenerateResponseTagsAppliedBy) IsKnown() bool {
	switch r {
	case ThreatSignalArticleTagGenerateResponseTagsAppliedByAI, ThreatSignalArticleTagGenerateResponseTagsAppliedByAnalyst, ThreatSignalArticleTagGenerateResponseTagsAppliedBySystem:
		return true
	}
	return false
}

type ThreatSignalArticleTagNewParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	TagID     param.Field[string] `json:"tag_id" api:"required" format:"uuid"`
}

func (r ThreatSignalArticleTagNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ThreatSignalArticleTagNewResponseEnvelope struct {
	Errors   []ThreatSignalArticleTagNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ThreatSignalArticleTagNewResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   ThreatSignalArticleTagNewResponse                   `json:"result" api:"required"`
	Success  ThreatSignalArticleTagNewResponseEnvelopeSuccess    `json:"success" api:"required"`
	JSON     threatSignalArticleTagNewResponseEnvelopeJSON       `json:"-"`
}

// threatSignalArticleTagNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [ThreatSignalArticleTagNewResponseEnvelope]
type threatSignalArticleTagNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalArticleTagNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalArticleTagNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalArticleTagNewResponseEnvelopeErrors struct {
	Message string                                              `json:"message" api:"required"`
	JSON    threatSignalArticleTagNewResponseEnvelopeErrorsJSON `json:"-"`
}

// threatSignalArticleTagNewResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ThreatSignalArticleTagNewResponseEnvelopeErrors]
type threatSignalArticleTagNewResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalArticleTagNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalArticleTagNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalArticleTagNewResponseEnvelopeMessages struct {
	Message  string                                                  `json:"message" api:"required"`
	Code     float64                                                 `json:"code"`
	Expected string                                                  `json:"expected"`
	Path     []string                                                `json:"path"`
	Reason   ThreatSignalArticleTagNewResponseEnvelopeMessagesReason `json:"reason"`
	JSON     threatSignalArticleTagNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// threatSignalArticleTagNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ThreatSignalArticleTagNewResponseEnvelopeMessages]
type threatSignalArticleTagNewResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Expected    apijson.Field
	Path        apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalArticleTagNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalArticleTagNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalArticleTagNewResponseEnvelopeMessagesReason string

const (
	ThreatSignalArticleTagNewResponseEnvelopeMessagesReasonFreeCustomFeedLimit              ThreatSignalArticleTagNewResponseEnvelopeMessagesReason = "free_custom_feed_limit"
	ThreatSignalArticleTagNewResponseEnvelopeMessagesReasonFreeCustomSkillsDisabled         ThreatSignalArticleTagNewResponseEnvelopeMessagesReason = "free_custom_skills_disabled"
	ThreatSignalArticleTagNewResponseEnvelopeMessagesReasonFreeTierReconciliationInProgress ThreatSignalArticleTagNewResponseEnvelopeMessagesReason = "free_tier_reconciliation_in_progress"
	ThreatSignalArticleTagNewResponseEnvelopeMessagesReasonFreeTierReconciliationFailed     ThreatSignalArticleTagNewResponseEnvelopeMessagesReason = "free_tier_reconciliation_failed"
	ThreatSignalArticleTagNewResponseEnvelopeMessagesReasonRawContentExpired                ThreatSignalArticleTagNewResponseEnvelopeMessagesReason = "raw_content_expired"
	ThreatSignalArticleTagNewResponseEnvelopeMessagesReasonCuratedVisibilityChanged         ThreatSignalArticleTagNewResponseEnvelopeMessagesReason = "curated_visibility_changed"
)

func (r ThreatSignalArticleTagNewResponseEnvelopeMessagesReason) IsKnown() bool {
	switch r {
	case ThreatSignalArticleTagNewResponseEnvelopeMessagesReasonFreeCustomFeedLimit, ThreatSignalArticleTagNewResponseEnvelopeMessagesReasonFreeCustomSkillsDisabled, ThreatSignalArticleTagNewResponseEnvelopeMessagesReasonFreeTierReconciliationInProgress, ThreatSignalArticleTagNewResponseEnvelopeMessagesReasonFreeTierReconciliationFailed, ThreatSignalArticleTagNewResponseEnvelopeMessagesReasonRawContentExpired, ThreatSignalArticleTagNewResponseEnvelopeMessagesReasonCuratedVisibilityChanged:
		return true
	}
	return false
}

type ThreatSignalArticleTagNewResponseEnvelopeSuccess bool

const (
	ThreatSignalArticleTagNewResponseEnvelopeSuccessTrue ThreatSignalArticleTagNewResponseEnvelopeSuccess = true
)

func (r ThreatSignalArticleTagNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ThreatSignalArticleTagNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ThreatSignalArticleTagDeleteParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type ThreatSignalArticleTagDeleteResponseEnvelope struct {
	Errors   []ThreatSignalArticleTagDeleteResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ThreatSignalArticleTagDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   ThreatSignalArticleTagDeleteResponse                   `json:"result" api:"required"`
	Success  ThreatSignalArticleTagDeleteResponseEnvelopeSuccess    `json:"success" api:"required"`
	JSON     threatSignalArticleTagDeleteResponseEnvelopeJSON       `json:"-"`
}

// threatSignalArticleTagDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [ThreatSignalArticleTagDeleteResponseEnvelope]
type threatSignalArticleTagDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalArticleTagDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalArticleTagDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalArticleTagDeleteResponseEnvelopeErrors struct {
	Message string                                                 `json:"message" api:"required"`
	JSON    threatSignalArticleTagDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// threatSignalArticleTagDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ThreatSignalArticleTagDeleteResponseEnvelopeErrors]
type threatSignalArticleTagDeleteResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalArticleTagDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalArticleTagDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalArticleTagDeleteResponseEnvelopeMessages struct {
	Message  string                                                     `json:"message" api:"required"`
	Code     float64                                                    `json:"code"`
	Expected string                                                     `json:"expected"`
	Path     []string                                                   `json:"path"`
	Reason   ThreatSignalArticleTagDeleteResponseEnvelopeMessagesReason `json:"reason"`
	JSON     threatSignalArticleTagDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// threatSignalArticleTagDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ThreatSignalArticleTagDeleteResponseEnvelopeMessages]
type threatSignalArticleTagDeleteResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Expected    apijson.Field
	Path        apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalArticleTagDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalArticleTagDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalArticleTagDeleteResponseEnvelopeMessagesReason string

const (
	ThreatSignalArticleTagDeleteResponseEnvelopeMessagesReasonFreeCustomFeedLimit              ThreatSignalArticleTagDeleteResponseEnvelopeMessagesReason = "free_custom_feed_limit"
	ThreatSignalArticleTagDeleteResponseEnvelopeMessagesReasonFreeCustomSkillsDisabled         ThreatSignalArticleTagDeleteResponseEnvelopeMessagesReason = "free_custom_skills_disabled"
	ThreatSignalArticleTagDeleteResponseEnvelopeMessagesReasonFreeTierReconciliationInProgress ThreatSignalArticleTagDeleteResponseEnvelopeMessagesReason = "free_tier_reconciliation_in_progress"
	ThreatSignalArticleTagDeleteResponseEnvelopeMessagesReasonFreeTierReconciliationFailed     ThreatSignalArticleTagDeleteResponseEnvelopeMessagesReason = "free_tier_reconciliation_failed"
	ThreatSignalArticleTagDeleteResponseEnvelopeMessagesReasonRawContentExpired                ThreatSignalArticleTagDeleteResponseEnvelopeMessagesReason = "raw_content_expired"
	ThreatSignalArticleTagDeleteResponseEnvelopeMessagesReasonCuratedVisibilityChanged         ThreatSignalArticleTagDeleteResponseEnvelopeMessagesReason = "curated_visibility_changed"
)

func (r ThreatSignalArticleTagDeleteResponseEnvelopeMessagesReason) IsKnown() bool {
	switch r {
	case ThreatSignalArticleTagDeleteResponseEnvelopeMessagesReasonFreeCustomFeedLimit, ThreatSignalArticleTagDeleteResponseEnvelopeMessagesReasonFreeCustomSkillsDisabled, ThreatSignalArticleTagDeleteResponseEnvelopeMessagesReasonFreeTierReconciliationInProgress, ThreatSignalArticleTagDeleteResponseEnvelopeMessagesReasonFreeTierReconciliationFailed, ThreatSignalArticleTagDeleteResponseEnvelopeMessagesReasonRawContentExpired, ThreatSignalArticleTagDeleteResponseEnvelopeMessagesReasonCuratedVisibilityChanged:
		return true
	}
	return false
}

type ThreatSignalArticleTagDeleteResponseEnvelopeSuccess bool

const (
	ThreatSignalArticleTagDeleteResponseEnvelopeSuccessTrue ThreatSignalArticleTagDeleteResponseEnvelopeSuccess = true
)

func (r ThreatSignalArticleTagDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ThreatSignalArticleTagDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ThreatSignalArticleTagGenerateParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type ThreatSignalArticleTagGenerateResponseEnvelope struct {
	Errors   []ThreatSignalArticleTagGenerateResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ThreatSignalArticleTagGenerateResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   ThreatSignalArticleTagGenerateResponse                   `json:"result" api:"required"`
	Success  ThreatSignalArticleTagGenerateResponseEnvelopeSuccess    `json:"success" api:"required"`
	JSON     threatSignalArticleTagGenerateResponseEnvelopeJSON       `json:"-"`
}

// threatSignalArticleTagGenerateResponseEnvelopeJSON contains the JSON metadata
// for the struct [ThreatSignalArticleTagGenerateResponseEnvelope]
type threatSignalArticleTagGenerateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalArticleTagGenerateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalArticleTagGenerateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalArticleTagGenerateResponseEnvelopeErrors struct {
	Message string                                                   `json:"message" api:"required"`
	JSON    threatSignalArticleTagGenerateResponseEnvelopeErrorsJSON `json:"-"`
}

// threatSignalArticleTagGenerateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ThreatSignalArticleTagGenerateResponseEnvelopeErrors]
type threatSignalArticleTagGenerateResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalArticleTagGenerateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalArticleTagGenerateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalArticleTagGenerateResponseEnvelopeMessages struct {
	Message  string                                                       `json:"message" api:"required"`
	Code     float64                                                      `json:"code"`
	Expected string                                                       `json:"expected"`
	Path     []string                                                     `json:"path"`
	Reason   ThreatSignalArticleTagGenerateResponseEnvelopeMessagesReason `json:"reason"`
	JSON     threatSignalArticleTagGenerateResponseEnvelopeMessagesJSON   `json:"-"`
}

// threatSignalArticleTagGenerateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ThreatSignalArticleTagGenerateResponseEnvelopeMessages]
type threatSignalArticleTagGenerateResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Expected    apijson.Field
	Path        apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalArticleTagGenerateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalArticleTagGenerateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalArticleTagGenerateResponseEnvelopeMessagesReason string

const (
	ThreatSignalArticleTagGenerateResponseEnvelopeMessagesReasonFreeCustomFeedLimit              ThreatSignalArticleTagGenerateResponseEnvelopeMessagesReason = "free_custom_feed_limit"
	ThreatSignalArticleTagGenerateResponseEnvelopeMessagesReasonFreeCustomSkillsDisabled         ThreatSignalArticleTagGenerateResponseEnvelopeMessagesReason = "free_custom_skills_disabled"
	ThreatSignalArticleTagGenerateResponseEnvelopeMessagesReasonFreeTierReconciliationInProgress ThreatSignalArticleTagGenerateResponseEnvelopeMessagesReason = "free_tier_reconciliation_in_progress"
	ThreatSignalArticleTagGenerateResponseEnvelopeMessagesReasonFreeTierReconciliationFailed     ThreatSignalArticleTagGenerateResponseEnvelopeMessagesReason = "free_tier_reconciliation_failed"
	ThreatSignalArticleTagGenerateResponseEnvelopeMessagesReasonRawContentExpired                ThreatSignalArticleTagGenerateResponseEnvelopeMessagesReason = "raw_content_expired"
	ThreatSignalArticleTagGenerateResponseEnvelopeMessagesReasonCuratedVisibilityChanged         ThreatSignalArticleTagGenerateResponseEnvelopeMessagesReason = "curated_visibility_changed"
)

func (r ThreatSignalArticleTagGenerateResponseEnvelopeMessagesReason) IsKnown() bool {
	switch r {
	case ThreatSignalArticleTagGenerateResponseEnvelopeMessagesReasonFreeCustomFeedLimit, ThreatSignalArticleTagGenerateResponseEnvelopeMessagesReasonFreeCustomSkillsDisabled, ThreatSignalArticleTagGenerateResponseEnvelopeMessagesReasonFreeTierReconciliationInProgress, ThreatSignalArticleTagGenerateResponseEnvelopeMessagesReasonFreeTierReconciliationFailed, ThreatSignalArticleTagGenerateResponseEnvelopeMessagesReasonRawContentExpired, ThreatSignalArticleTagGenerateResponseEnvelopeMessagesReasonCuratedVisibilityChanged:
		return true
	}
	return false
}

type ThreatSignalArticleTagGenerateResponseEnvelopeSuccess bool

const (
	ThreatSignalArticleTagGenerateResponseEnvelopeSuccessTrue ThreatSignalArticleTagGenerateResponseEnvelopeSuccess = true
)

func (r ThreatSignalArticleTagGenerateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ThreatSignalArticleTagGenerateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
