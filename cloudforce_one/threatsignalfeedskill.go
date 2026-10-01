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

// ThreatSignalFeedSkillService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreatSignalFeedSkillService] method instead.
type ThreatSignalFeedSkillService struct {
	Options []option.RequestOption
}

// NewThreatSignalFeedSkillService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewThreatSignalFeedSkillService(opts ...option.RequestOption) (r *ThreatSignalFeedSkillService) {
	r = &ThreatSignalFeedSkillService{}
	r.Options = opts
	return
}

// Replaces the ordered custom skills assigned to a Threat Signals feed.
func (r *ThreatSignalFeedSkillService) Update(ctx context.Context, feedID string, params ThreatSignalFeedSkillUpdateParams, opts ...option.RequestOption) (res *ThreatSignalFeedSkillUpdateResponse, err error) {
	var env ThreatSignalFeedSkillUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if feedID == "" {
		err = errors.New("missing required feed_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/v2/threat-signals/feeds/%s/skills", params.AccountID, feedID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieves the effective skill pipeline for a Threat Signals feed.
func (r *ThreatSignalFeedSkillService) Get(ctx context.Context, feedID string, query ThreatSignalFeedSkillGetParams, opts ...option.RequestOption) (res *ThreatSignalFeedSkillGetResponse, err error) {
	var env ThreatSignalFeedSkillGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if feedID == "" {
		err = errors.New("missing required feed_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/v2/threat-signals/feeds/%s/skills", query.AccountID, feedID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type ThreatSignalFeedSkillUpdateResponse struct {
	FeedID string                                     `json:"feed_id" api:"required" format:"uuid"`
	Skills []ThreatSignalFeedSkillUpdateResponseSkill `json:"skills" api:"required"`
	JSON   threatSignalFeedSkillUpdateResponseJSON    `json:"-"`
}

// threatSignalFeedSkillUpdateResponseJSON contains the JSON metadata for the
// struct [ThreatSignalFeedSkillUpdateResponse]
type threatSignalFeedSkillUpdateResponseJSON struct {
	FeedID      apijson.Field
	Skills      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalFeedSkillUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalFeedSkillUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalFeedSkillUpdateResponseSkill struct {
	// Zero-based pipeline position.
	Position int64                                        `json:"position" api:"required"`
	SkillID  string                                       `json:"skill_id" api:"required"`
	JSON     threatSignalFeedSkillUpdateResponseSkillJSON `json:"-"`
}

// threatSignalFeedSkillUpdateResponseSkillJSON contains the JSON metadata for the
// struct [ThreatSignalFeedSkillUpdateResponseSkill]
type threatSignalFeedSkillUpdateResponseSkillJSON struct {
	Position    apijson.Field
	SkillID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalFeedSkillUpdateResponseSkill) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalFeedSkillUpdateResponseSkillJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalFeedSkillGetResponse struct {
	FeedID string                                  `json:"feed_id" api:"required" format:"uuid"`
	Skills []ThreatSignalFeedSkillGetResponseSkill `json:"skills" api:"required"`
	JSON   threatSignalFeedSkillGetResponseJSON    `json:"-"`
}

// threatSignalFeedSkillGetResponseJSON contains the JSON metadata for the struct
// [ThreatSignalFeedSkillGetResponse]
type threatSignalFeedSkillGetResponseJSON struct {
	FeedID      apijson.Field
	Skills      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalFeedSkillGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalFeedSkillGetResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalFeedSkillGetResponseSkill struct {
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
	Source    ThreatSignalFeedSkillGetResponseSkillsSource `json:"source" api:"required"`
	Type      string                                       `json:"type" api:"required"`
	UpdatedAt string                                       `json:"updated_at" api:"required"`
	JSON      threatSignalFeedSkillGetResponseSkillJSON    `json:"-"`
}

// threatSignalFeedSkillGetResponseSkillJSON contains the JSON metadata for the
// struct [ThreatSignalFeedSkillGetResponseSkill]
type threatSignalFeedSkillGetResponseSkillJSON struct {
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

func (r *ThreatSignalFeedSkillGetResponseSkill) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalFeedSkillGetResponseSkillJSON) RawJSON() string {
	return r.raw
}

// `default` for Cloudforce One managed skills (read-only), `custom` for account
// skills.
type ThreatSignalFeedSkillGetResponseSkillsSource string

const (
	ThreatSignalFeedSkillGetResponseSkillsSourceDefault ThreatSignalFeedSkillGetResponseSkillsSource = "default"
	ThreatSignalFeedSkillGetResponseSkillsSourceCustom  ThreatSignalFeedSkillGetResponseSkillsSource = "custom"
)

func (r ThreatSignalFeedSkillGetResponseSkillsSource) IsKnown() bool {
	switch r {
	case ThreatSignalFeedSkillGetResponseSkillsSourceDefault, ThreatSignalFeedSkillGetResponseSkillsSourceCustom:
		return true
	}
	return false
}

type ThreatSignalFeedSkillUpdateParams struct {
	AccountID param.Field[string]   `path:"account_id" api:"required"`
	SkillIDs  param.Field[[]string] `json:"skill_ids" api:"required" format:"uuid"`
}

func (r ThreatSignalFeedSkillUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ThreatSignalFeedSkillUpdateResponseEnvelope struct {
	Errors   []ThreatSignalFeedSkillUpdateResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ThreatSignalFeedSkillUpdateResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   ThreatSignalFeedSkillUpdateResponse                   `json:"result" api:"required"`
	Success  ThreatSignalFeedSkillUpdateResponseEnvelopeSuccess    `json:"success" api:"required"`
	JSON     threatSignalFeedSkillUpdateResponseEnvelopeJSON       `json:"-"`
}

// threatSignalFeedSkillUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [ThreatSignalFeedSkillUpdateResponseEnvelope]
type threatSignalFeedSkillUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalFeedSkillUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalFeedSkillUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalFeedSkillUpdateResponseEnvelopeErrors struct {
	Message string                                                `json:"message" api:"required"`
	JSON    threatSignalFeedSkillUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// threatSignalFeedSkillUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ThreatSignalFeedSkillUpdateResponseEnvelopeErrors]
type threatSignalFeedSkillUpdateResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalFeedSkillUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalFeedSkillUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalFeedSkillUpdateResponseEnvelopeMessages struct {
	Message  string                                                    `json:"message" api:"required"`
	Code     float64                                                   `json:"code"`
	Expected string                                                    `json:"expected"`
	Path     []string                                                  `json:"path"`
	Reason   ThreatSignalFeedSkillUpdateResponseEnvelopeMessagesReason `json:"reason"`
	JSON     threatSignalFeedSkillUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// threatSignalFeedSkillUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ThreatSignalFeedSkillUpdateResponseEnvelopeMessages]
type threatSignalFeedSkillUpdateResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Expected    apijson.Field
	Path        apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalFeedSkillUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalFeedSkillUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalFeedSkillUpdateResponseEnvelopeMessagesReason string

const (
	ThreatSignalFeedSkillUpdateResponseEnvelopeMessagesReasonFreeCustomFeedLimit              ThreatSignalFeedSkillUpdateResponseEnvelopeMessagesReason = "free_custom_feed_limit"
	ThreatSignalFeedSkillUpdateResponseEnvelopeMessagesReasonFreeCustomSkillsDisabled         ThreatSignalFeedSkillUpdateResponseEnvelopeMessagesReason = "free_custom_skills_disabled"
	ThreatSignalFeedSkillUpdateResponseEnvelopeMessagesReasonFreeTierReconciliationInProgress ThreatSignalFeedSkillUpdateResponseEnvelopeMessagesReason = "free_tier_reconciliation_in_progress"
	ThreatSignalFeedSkillUpdateResponseEnvelopeMessagesReasonFreeTierReconciliationFailed     ThreatSignalFeedSkillUpdateResponseEnvelopeMessagesReason = "free_tier_reconciliation_failed"
	ThreatSignalFeedSkillUpdateResponseEnvelopeMessagesReasonRawContentExpired                ThreatSignalFeedSkillUpdateResponseEnvelopeMessagesReason = "raw_content_expired"
	ThreatSignalFeedSkillUpdateResponseEnvelopeMessagesReasonCuratedVisibilityChanged         ThreatSignalFeedSkillUpdateResponseEnvelopeMessagesReason = "curated_visibility_changed"
)

func (r ThreatSignalFeedSkillUpdateResponseEnvelopeMessagesReason) IsKnown() bool {
	switch r {
	case ThreatSignalFeedSkillUpdateResponseEnvelopeMessagesReasonFreeCustomFeedLimit, ThreatSignalFeedSkillUpdateResponseEnvelopeMessagesReasonFreeCustomSkillsDisabled, ThreatSignalFeedSkillUpdateResponseEnvelopeMessagesReasonFreeTierReconciliationInProgress, ThreatSignalFeedSkillUpdateResponseEnvelopeMessagesReasonFreeTierReconciliationFailed, ThreatSignalFeedSkillUpdateResponseEnvelopeMessagesReasonRawContentExpired, ThreatSignalFeedSkillUpdateResponseEnvelopeMessagesReasonCuratedVisibilityChanged:
		return true
	}
	return false
}

type ThreatSignalFeedSkillUpdateResponseEnvelopeSuccess bool

const (
	ThreatSignalFeedSkillUpdateResponseEnvelopeSuccessTrue ThreatSignalFeedSkillUpdateResponseEnvelopeSuccess = true
)

func (r ThreatSignalFeedSkillUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ThreatSignalFeedSkillUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ThreatSignalFeedSkillGetParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type ThreatSignalFeedSkillGetResponseEnvelope struct {
	Errors   []ThreatSignalFeedSkillGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ThreatSignalFeedSkillGetResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   ThreatSignalFeedSkillGetResponse                   `json:"result" api:"required"`
	Success  ThreatSignalFeedSkillGetResponseEnvelopeSuccess    `json:"success" api:"required"`
	JSON     threatSignalFeedSkillGetResponseEnvelopeJSON       `json:"-"`
}

// threatSignalFeedSkillGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ThreatSignalFeedSkillGetResponseEnvelope]
type threatSignalFeedSkillGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalFeedSkillGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalFeedSkillGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalFeedSkillGetResponseEnvelopeErrors struct {
	Message string                                             `json:"message" api:"required"`
	JSON    threatSignalFeedSkillGetResponseEnvelopeErrorsJSON `json:"-"`
}

// threatSignalFeedSkillGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ThreatSignalFeedSkillGetResponseEnvelopeErrors]
type threatSignalFeedSkillGetResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalFeedSkillGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalFeedSkillGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalFeedSkillGetResponseEnvelopeMessages struct {
	Message  string                                                 `json:"message" api:"required"`
	Code     float64                                                `json:"code"`
	Expected string                                                 `json:"expected"`
	Path     []string                                               `json:"path"`
	Reason   ThreatSignalFeedSkillGetResponseEnvelopeMessagesReason `json:"reason"`
	JSON     threatSignalFeedSkillGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// threatSignalFeedSkillGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ThreatSignalFeedSkillGetResponseEnvelopeMessages]
type threatSignalFeedSkillGetResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Expected    apijson.Field
	Path        apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalFeedSkillGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalFeedSkillGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalFeedSkillGetResponseEnvelopeMessagesReason string

const (
	ThreatSignalFeedSkillGetResponseEnvelopeMessagesReasonFreeCustomFeedLimit              ThreatSignalFeedSkillGetResponseEnvelopeMessagesReason = "free_custom_feed_limit"
	ThreatSignalFeedSkillGetResponseEnvelopeMessagesReasonFreeCustomSkillsDisabled         ThreatSignalFeedSkillGetResponseEnvelopeMessagesReason = "free_custom_skills_disabled"
	ThreatSignalFeedSkillGetResponseEnvelopeMessagesReasonFreeTierReconciliationInProgress ThreatSignalFeedSkillGetResponseEnvelopeMessagesReason = "free_tier_reconciliation_in_progress"
	ThreatSignalFeedSkillGetResponseEnvelopeMessagesReasonFreeTierReconciliationFailed     ThreatSignalFeedSkillGetResponseEnvelopeMessagesReason = "free_tier_reconciliation_failed"
	ThreatSignalFeedSkillGetResponseEnvelopeMessagesReasonRawContentExpired                ThreatSignalFeedSkillGetResponseEnvelopeMessagesReason = "raw_content_expired"
	ThreatSignalFeedSkillGetResponseEnvelopeMessagesReasonCuratedVisibilityChanged         ThreatSignalFeedSkillGetResponseEnvelopeMessagesReason = "curated_visibility_changed"
)

func (r ThreatSignalFeedSkillGetResponseEnvelopeMessagesReason) IsKnown() bool {
	switch r {
	case ThreatSignalFeedSkillGetResponseEnvelopeMessagesReasonFreeCustomFeedLimit, ThreatSignalFeedSkillGetResponseEnvelopeMessagesReasonFreeCustomSkillsDisabled, ThreatSignalFeedSkillGetResponseEnvelopeMessagesReasonFreeTierReconciliationInProgress, ThreatSignalFeedSkillGetResponseEnvelopeMessagesReasonFreeTierReconciliationFailed, ThreatSignalFeedSkillGetResponseEnvelopeMessagesReasonRawContentExpired, ThreatSignalFeedSkillGetResponseEnvelopeMessagesReasonCuratedVisibilityChanged:
		return true
	}
	return false
}

type ThreatSignalFeedSkillGetResponseEnvelopeSuccess bool

const (
	ThreatSignalFeedSkillGetResponseEnvelopeSuccessTrue ThreatSignalFeedSkillGetResponseEnvelopeSuccess = true
)

func (r ThreatSignalFeedSkillGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ThreatSignalFeedSkillGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
