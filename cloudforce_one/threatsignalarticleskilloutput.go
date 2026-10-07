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

// ThreatSignalArticleSkillOutputService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreatSignalArticleSkillOutputService] method instead.
type ThreatSignalArticleSkillOutputService struct {
	Options []option.RequestOption
}

// NewThreatSignalArticleSkillOutputService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewThreatSignalArticleSkillOutputService(opts ...option.RequestOption) (r *ThreatSignalArticleSkillOutputService) {
	r = &ThreatSignalArticleSkillOutputService{}
	r.Options = opts
	return
}

// Retrieves the stored output of a skill for a Threat Signals article.
func (r *ThreatSignalArticleSkillOutputService) Get(ctx context.Context, articleID string, skillID string, query ThreatSignalArticleSkillOutputGetParams, opts ...option.RequestOption) (res *ThreatSignalArticleSkillOutputGetResponse, err error) {
	var env ThreatSignalArticleSkillOutputGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if articleID == "" {
		err = errors.New("missing required article_id parameter")
		return nil, err
	}
	if skillID == "" {
		err = errors.New("missing required skill_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/v2/threat-signals/articles/%s/skills/%s/output", query.AccountID, articleID, skillID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type ThreatSignalArticleSkillOutputGetResponse struct {
	ArticleID          string `json:"article_id" api:"required" format:"uuid"`
	CustomSkillVersion string `json:"custom_skill_version" api:"required,nullable"`
	// JSON-encoded output schema of the skill. Null when the skill no longer exists.
	OutputSchema string `json:"output_schema" api:"required,nullable"`
	SkillID      string `json:"skill_id" api:"required"`
	// Skill output. Parsed JSON when the stored output is valid JSON, otherwise the
	// raw string.
	CustomOutput interface{}                                   `json:"custom_output"`
	JSON         threatSignalArticleSkillOutputGetResponseJSON `json:"-"`
}

// threatSignalArticleSkillOutputGetResponseJSON contains the JSON metadata for the
// struct [ThreatSignalArticleSkillOutputGetResponse]
type threatSignalArticleSkillOutputGetResponseJSON struct {
	ArticleID          apijson.Field
	CustomSkillVersion apijson.Field
	OutputSchema       apijson.Field
	SkillID            apijson.Field
	CustomOutput       apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ThreatSignalArticleSkillOutputGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalArticleSkillOutputGetResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalArticleSkillOutputGetParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type ThreatSignalArticleSkillOutputGetResponseEnvelope struct {
	Errors   []ThreatSignalArticleSkillOutputGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ThreatSignalArticleSkillOutputGetResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   ThreatSignalArticleSkillOutputGetResponse                   `json:"result" api:"required"`
	Success  ThreatSignalArticleSkillOutputGetResponseEnvelopeSuccess    `json:"success" api:"required"`
	JSON     threatSignalArticleSkillOutputGetResponseEnvelopeJSON       `json:"-"`
}

// threatSignalArticleSkillOutputGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [ThreatSignalArticleSkillOutputGetResponseEnvelope]
type threatSignalArticleSkillOutputGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalArticleSkillOutputGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalArticleSkillOutputGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalArticleSkillOutputGetResponseEnvelopeErrors struct {
	Message string                                                      `json:"message" api:"required"`
	JSON    threatSignalArticleSkillOutputGetResponseEnvelopeErrorsJSON `json:"-"`
}

// threatSignalArticleSkillOutputGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [ThreatSignalArticleSkillOutputGetResponseEnvelopeErrors]
type threatSignalArticleSkillOutputGetResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalArticleSkillOutputGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalArticleSkillOutputGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalArticleSkillOutputGetResponseEnvelopeMessages struct {
	Message  string                                                          `json:"message" api:"required"`
	Code     float64                                                         `json:"code"`
	Expected string                                                          `json:"expected"`
	Path     []string                                                        `json:"path"`
	Reason   ThreatSignalArticleSkillOutputGetResponseEnvelopeMessagesReason `json:"reason"`
	JSON     threatSignalArticleSkillOutputGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// threatSignalArticleSkillOutputGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ThreatSignalArticleSkillOutputGetResponseEnvelopeMessages]
type threatSignalArticleSkillOutputGetResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Expected    apijson.Field
	Path        apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalArticleSkillOutputGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalArticleSkillOutputGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalArticleSkillOutputGetResponseEnvelopeMessagesReason string

const (
	ThreatSignalArticleSkillOutputGetResponseEnvelopeMessagesReasonFreeCustomFeedLimit              ThreatSignalArticleSkillOutputGetResponseEnvelopeMessagesReason = "free_custom_feed_limit"
	ThreatSignalArticleSkillOutputGetResponseEnvelopeMessagesReasonFreeCustomSkillsDisabled         ThreatSignalArticleSkillOutputGetResponseEnvelopeMessagesReason = "free_custom_skills_disabled"
	ThreatSignalArticleSkillOutputGetResponseEnvelopeMessagesReasonFreeTierReconciliationInProgress ThreatSignalArticleSkillOutputGetResponseEnvelopeMessagesReason = "free_tier_reconciliation_in_progress"
	ThreatSignalArticleSkillOutputGetResponseEnvelopeMessagesReasonFreeTierReconciliationFailed     ThreatSignalArticleSkillOutputGetResponseEnvelopeMessagesReason = "free_tier_reconciliation_failed"
	ThreatSignalArticleSkillOutputGetResponseEnvelopeMessagesReasonRawContentExpired                ThreatSignalArticleSkillOutputGetResponseEnvelopeMessagesReason = "raw_content_expired"
	ThreatSignalArticleSkillOutputGetResponseEnvelopeMessagesReasonCuratedVisibilityChanged         ThreatSignalArticleSkillOutputGetResponseEnvelopeMessagesReason = "curated_visibility_changed"
)

func (r ThreatSignalArticleSkillOutputGetResponseEnvelopeMessagesReason) IsKnown() bool {
	switch r {
	case ThreatSignalArticleSkillOutputGetResponseEnvelopeMessagesReasonFreeCustomFeedLimit, ThreatSignalArticleSkillOutputGetResponseEnvelopeMessagesReasonFreeCustomSkillsDisabled, ThreatSignalArticleSkillOutputGetResponseEnvelopeMessagesReasonFreeTierReconciliationInProgress, ThreatSignalArticleSkillOutputGetResponseEnvelopeMessagesReasonFreeTierReconciliationFailed, ThreatSignalArticleSkillOutputGetResponseEnvelopeMessagesReasonRawContentExpired, ThreatSignalArticleSkillOutputGetResponseEnvelopeMessagesReasonCuratedVisibilityChanged:
		return true
	}
	return false
}

type ThreatSignalArticleSkillOutputGetResponseEnvelopeSuccess bool

const (
	ThreatSignalArticleSkillOutputGetResponseEnvelopeSuccessTrue ThreatSignalArticleSkillOutputGetResponseEnvelopeSuccess = true
)

func (r ThreatSignalArticleSkillOutputGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ThreatSignalArticleSkillOutputGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
