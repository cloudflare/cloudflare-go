// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/shared"
	"github.com/tidwall/gjson"
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

// Retrieves the stored output of a skill for a Threat Signals article. The
// custom_output field is parsed JSON when the stored output is valid JSON,
// otherwise the raw string.
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
	// Article UUID that received the skill output.
	ArticleID string `json:"article_id" api:"required" format:"uuid"`
	// Untrusted model output: parsed JSON when the stored completion text is valid
	// JSON, otherwise raw text. Consumers must safely render or escape it.
	CustomOutput ThreatSignalArticleSkillOutputGetResponseCustomOutputUnion `json:"custom_output" api:"required"`
	// Whether custom_output was parsed from the stored completion text.
	CustomOutputParseStatus ThreatSignalArticleSkillOutputGetResponseCustomOutputParseStatus `json:"custom_output_parse_status" api:"required"`
	// Non-blocking validation classification for untrusted model output. `unknown` is
	// retained only for historical rows.
	CustomOutputValidationStatus ThreatSignalArticleSkillOutputGetResponseCustomOutputValidationStatus `json:"custom_output_validation_status" api:"required"`
	// Custom Skill version that produced the output. Null when historical metadata is
	// unavailable.
	CustomSkillVersion string `json:"custom_skill_version" api:"required,nullable"`
	// JSON-encoded output schema. Null when historical skill metadata is unavailable.
	OutputSchema string `json:"output_schema" api:"required,nullable"`
	// Custom Skill UUID that produced the output.
	SkillID string                                        `json:"skill_id" api:"required"`
	JSON    threatSignalArticleSkillOutputGetResponseJSON `json:"-"`
}

// threatSignalArticleSkillOutputGetResponseJSON contains the JSON metadata for the
// struct [ThreatSignalArticleSkillOutputGetResponse]
type threatSignalArticleSkillOutputGetResponseJSON struct {
	ArticleID                    apijson.Field
	CustomOutput                 apijson.Field
	CustomOutputParseStatus      apijson.Field
	CustomOutputValidationStatus apijson.Field
	CustomSkillVersion           apijson.Field
	OutputSchema                 apijson.Field
	SkillID                      apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *ThreatSignalArticleSkillOutputGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalArticleSkillOutputGetResponseJSON) RawJSON() string {
	return r.raw
}

// Untrusted model output: parsed JSON when the stored completion text is valid
// JSON, otherwise raw text. Consumers must safely render or escape it.
//
// Union satisfied by [shared.UnionString], [shared.UnionFloat],
// [shared.UnionBool], [ThreatSignalArticleSkillOutputGetResponseCustomOutputMap]
// or [ThreatSignalArticleSkillOutputGetResponseCustomOutputArray].
type ThreatSignalArticleSkillOutputGetResponseCustomOutputUnion interface {
	ImplementsThreatSignalArticleSkillOutputGetResponseCustomOutputUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ThreatSignalArticleSkillOutputGetResponseCustomOutputUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ThreatSignalArticleSkillOutputGetResponseCustomOutputMap{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ThreatSignalArticleSkillOutputGetResponseCustomOutputArray{}),
		},
	)
}

type ThreatSignalArticleSkillOutputGetResponseCustomOutputMap map[string]interface{}

func (r ThreatSignalArticleSkillOutputGetResponseCustomOutputMap) ImplementsThreatSignalArticleSkillOutputGetResponseCustomOutputUnion() {
}

type ThreatSignalArticleSkillOutputGetResponseCustomOutputArray []ThreatSignalArticleSkillOutputGetResponseCustomOutputArrayUnionItem

func (r ThreatSignalArticleSkillOutputGetResponseCustomOutputArray) ImplementsThreatSignalArticleSkillOutputGetResponseCustomOutputUnion() {
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool]
// or [ThreatSignalArticleSkillOutputGetResponseCustomOutputArrayMap].
type ThreatSignalArticleSkillOutputGetResponseCustomOutputArrayUnionItem interface {
	ImplementsThreatSignalArticleSkillOutputGetResponseCustomOutputArrayUnionItem()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ThreatSignalArticleSkillOutputGetResponseCustomOutputArrayUnionItem)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ThreatSignalArticleSkillOutputGetResponseCustomOutputArrayMap{}),
		},
	)
}

type ThreatSignalArticleSkillOutputGetResponseCustomOutputArrayMap map[string]interface{}

func (r ThreatSignalArticleSkillOutputGetResponseCustomOutputArrayMap) ImplementsThreatSignalArticleSkillOutputGetResponseCustomOutputArrayUnionItem() {
}

// Whether custom_output was parsed from the stored completion text.
type ThreatSignalArticleSkillOutputGetResponseCustomOutputParseStatus string

const (
	ThreatSignalArticleSkillOutputGetResponseCustomOutputParseStatusParsed      ThreatSignalArticleSkillOutputGetResponseCustomOutputParseStatus = "parsed"
	ThreatSignalArticleSkillOutputGetResponseCustomOutputParseStatusInvalidJson ThreatSignalArticleSkillOutputGetResponseCustomOutputParseStatus = "invalid_json"
)

func (r ThreatSignalArticleSkillOutputGetResponseCustomOutputParseStatus) IsKnown() bool {
	switch r {
	case ThreatSignalArticleSkillOutputGetResponseCustomOutputParseStatusParsed, ThreatSignalArticleSkillOutputGetResponseCustomOutputParseStatusInvalidJson:
		return true
	}
	return false
}

// Non-blocking validation classification for untrusted model output. `unknown` is
// retained only for historical rows.
type ThreatSignalArticleSkillOutputGetResponseCustomOutputValidationStatus string

const (
	ThreatSignalArticleSkillOutputGetResponseCustomOutputValidationStatusValid         ThreatSignalArticleSkillOutputGetResponseCustomOutputValidationStatus = "valid"
	ThreatSignalArticleSkillOutputGetResponseCustomOutputValidationStatusInvalidJson   ThreatSignalArticleSkillOutputGetResponseCustomOutputValidationStatus = "invalid_json"
	ThreatSignalArticleSkillOutputGetResponseCustomOutputValidationStatusSchemaInvalid ThreatSignalArticleSkillOutputGetResponseCustomOutputValidationStatus = "schema_invalid"
	ThreatSignalArticleSkillOutputGetResponseCustomOutputValidationStatusUnknown       ThreatSignalArticleSkillOutputGetResponseCustomOutputValidationStatus = "unknown"
)

func (r ThreatSignalArticleSkillOutputGetResponseCustomOutputValidationStatus) IsKnown() bool {
	switch r {
	case ThreatSignalArticleSkillOutputGetResponseCustomOutputValidationStatusValid, ThreatSignalArticleSkillOutputGetResponseCustomOutputValidationStatusInvalidJson, ThreatSignalArticleSkillOutputGetResponseCustomOutputValidationStatusSchemaInvalid, ThreatSignalArticleSkillOutputGetResponseCustomOutputValidationStatusUnknown:
		return true
	}
	return false
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
	Message string  `json:"message" api:"required"`
	Code    float64 `json:"code"`
	// The current count of custom feeds for the account.
	CustomFeedCount int64 `json:"custom_feed_count"`
	// The custom feed limit for the account.
	CustomFeedLimit int64                                                           `json:"custom_feed_limit"`
	Expected        string                                                          `json:"expected"`
	Path            []string                                                        `json:"path"`
	Reason          ThreatSignalArticleSkillOutputGetResponseEnvelopeMessagesReason `json:"reason"`
	JSON            threatSignalArticleSkillOutputGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// threatSignalArticleSkillOutputGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ThreatSignalArticleSkillOutputGetResponseEnvelopeMessages]
type threatSignalArticleSkillOutputGetResponseEnvelopeMessagesJSON struct {
	Message         apijson.Field
	Code            apijson.Field
	CustomFeedCount apijson.Field
	CustomFeedLimit apijson.Field
	Expected        apijson.Field
	Path            apijson.Field
	Reason          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ThreatSignalArticleSkillOutputGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalArticleSkillOutputGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalArticleSkillOutputGetResponseEnvelopeMessagesReason string

const (
	ThreatSignalArticleSkillOutputGetResponseEnvelopeMessagesReasonThreatSignalsFeedLimit           ThreatSignalArticleSkillOutputGetResponseEnvelopeMessagesReason = "threat_signals_feed_limit"
	ThreatSignalArticleSkillOutputGetResponseEnvelopeMessagesReasonFreeCustomSkillsDisabled         ThreatSignalArticleSkillOutputGetResponseEnvelopeMessagesReason = "free_custom_skills_disabled"
	ThreatSignalArticleSkillOutputGetResponseEnvelopeMessagesReasonFreeSkillRunDisabled             ThreatSignalArticleSkillOutputGetResponseEnvelopeMessagesReason = "free_skill_run_disabled"
	ThreatSignalArticleSkillOutputGetResponseEnvelopeMessagesReasonFreeTierReconciliationInProgress ThreatSignalArticleSkillOutputGetResponseEnvelopeMessagesReason = "free_tier_reconciliation_in_progress"
	ThreatSignalArticleSkillOutputGetResponseEnvelopeMessagesReasonFreeTierReconciliationFailed     ThreatSignalArticleSkillOutputGetResponseEnvelopeMessagesReason = "free_tier_reconciliation_failed"
	ThreatSignalArticleSkillOutputGetResponseEnvelopeMessagesReasonRawContentExpired                ThreatSignalArticleSkillOutputGetResponseEnvelopeMessagesReason = "raw_content_expired"
	ThreatSignalArticleSkillOutputGetResponseEnvelopeMessagesReasonCuratedVisibilityChanged         ThreatSignalArticleSkillOutputGetResponseEnvelopeMessagesReason = "curated_visibility_changed"
)

func (r ThreatSignalArticleSkillOutputGetResponseEnvelopeMessagesReason) IsKnown() bool {
	switch r {
	case ThreatSignalArticleSkillOutputGetResponseEnvelopeMessagesReasonThreatSignalsFeedLimit, ThreatSignalArticleSkillOutputGetResponseEnvelopeMessagesReasonFreeCustomSkillsDisabled, ThreatSignalArticleSkillOutputGetResponseEnvelopeMessagesReasonFreeSkillRunDisabled, ThreatSignalArticleSkillOutputGetResponseEnvelopeMessagesReasonFreeTierReconciliationInProgress, ThreatSignalArticleSkillOutputGetResponseEnvelopeMessagesReasonFreeTierReconciliationFailed, ThreatSignalArticleSkillOutputGetResponseEnvelopeMessagesReasonRawContentExpired, ThreatSignalArticleSkillOutputGetResponseEnvelopeMessagesReasonCuratedVisibilityChanged:
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
