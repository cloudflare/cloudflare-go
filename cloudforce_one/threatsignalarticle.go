// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one

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
)

// ThreatSignalArticleService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreatSignalArticleService] method instead.
type ThreatSignalArticleService struct {
	Options      []option.RequestOption
	Content      *ThreatSignalArticleContentService
	Tags         *ThreatSignalArticleTagService
	SkillOutputs *ThreatSignalArticleSkillOutputService
}

// NewThreatSignalArticleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewThreatSignalArticleService(opts ...option.RequestOption) (r *ThreatSignalArticleService) {
	r = &ThreatSignalArticleService{}
	r.Options = opts
	r.Content = NewThreatSignalArticleContentService(opts...)
	r.Tags = NewThreatSignalArticleTagService(opts...)
	r.SkillOutputs = NewThreatSignalArticleSkillOutputService(opts...)
	return
}

// Lists articles from the account's Threat Signals feeds.
func (r *ThreatSignalArticleService) List(ctx context.Context, params ThreatSignalArticleListParams, opts ...option.RequestOption) (res *ThreatSignalArticleListResponse, err error) {
	var env ThreatSignalArticleListResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/v2/threat-signals/articles", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Marks up to 50 Threat Signals articles as read or unread.
func (r *ThreatSignalArticleService) BulkEdit(ctx context.Context, params ThreatSignalArticleBulkEditParams, opts ...option.RequestOption) (res *ThreatSignalArticleBulkEditResponse, err error) {
	var env ThreatSignalArticleBulkEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/v2/threat-signals/articles", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Marks a Threat Signals article as read or unread.
func (r *ThreatSignalArticleService) Edit(ctx context.Context, articleID string, params ThreatSignalArticleEditParams, opts ...option.RequestOption) (res *ThreatSignalArticleEditResponse, err error) {
	var env ThreatSignalArticleEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if articleID == "" {
		err = errors.New("missing required article_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/v2/threat-signals/articles/%s", params.AccountID, articleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieves a Threat Signals article with its summary, tags, indicator status and
// relevance to the account's Priority Intelligence Requirements.
func (r *ThreatSignalArticleService) Get(ctx context.Context, articleID string, query ThreatSignalArticleGetParams, opts ...option.RequestOption) (res *ThreatSignalArticleGetResponse, err error) {
	var env ThreatSignalArticleGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if articleID == "" {
		err = errors.New("missing required article_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/v2/threat-signals/articles/%s", query.AccountID, articleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type ThreatSignalArticleListResponse struct {
	Articles          []ThreatSignalArticleListResponseArticle `json:"articles" api:"required"`
	HasMore           bool                                     `json:"has_more" api:"required"`
	NextCursor        string                                   `json:"next_cursor" api:"required,nullable"`
	TotalCount        float64                                  `json:"total_count" api:"required,nullable"`
	TotalCountIsExact bool                                     `json:"total_count_is_exact" api:"required"`
	JSON              threatSignalArticleListResponseJSON      `json:"-"`
}

// threatSignalArticleListResponseJSON contains the JSON metadata for the struct
// [ThreatSignalArticleListResponse]
type threatSignalArticleListResponseJSON struct {
	Articles          apijson.Field
	HasMore           apijson.Field
	NextCursor        apijson.Field
	TotalCount        apijson.Field
	TotalCountIsExact apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ThreatSignalArticleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalArticleListResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalArticleListResponseArticle struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Threat Events dataset identifier for the article redirect. Null when the account
	// feeds dataset mapping is unavailable.
	DatasetID string `json:"dataset_id" api:"required,nullable"`
	// Threat Events event identifier associated with this article for a UI redirect.
	// Null when no event has been linked.
	EventID         string `json:"event_id" api:"required,nullable"`
	FeedDisplayName string `json:"feed_display_name" api:"required,nullable"`
	FeedID          string `json:"feed_id" api:"required" format:"uuid"`
	FetchedAt       string `json:"fetched_at" api:"required"`
	Link            string `json:"link" api:"required,nullable"`
	PublishedAt     string `json:"published_at" api:"required,nullable"`
	Read            bool   `json:"read" api:"required"`
	ReadAt          string `json:"read_at" api:"required,nullable"`
	// Beta. The article's relevance to the account's active Priority Intelligence
	// Requirements. Null when the article has not been scored against them, including
	// when the account has no active PIRs, or has not been scored against the
	// requested pir_id.
	Relevance ThreatSignalArticleListResponseArticlesRelevance `json:"relevance" api:"required,nullable"`
	// Beta. Intrinsic article severity band, from informational to critical,
	// independent of any organisation or Priority Intelligence Requirement. Null when
	// the account has no active Priority Intelligence Requirements, scoring has not
	// completed, or a historic article has not entered an automatic rescore window.
	Severity ThreatSignalArticleListResponseArticlesSeverity `json:"severity" api:"required,nullable"`
	// Persisted enrichment summary. Null until enrichment produces a summary.
	Summary string                                       `json:"summary" api:"required,nullable"`
	Tags    []ThreatSignalArticleListResponseArticlesTag `json:"tags" api:"required"`
	Title   string                                       `json:"title" api:"required,nullable"`
	JSON    threatSignalArticleListResponseArticleJSON   `json:"-"`
}

// threatSignalArticleListResponseArticleJSON contains the JSON metadata for the
// struct [ThreatSignalArticleListResponseArticle]
type threatSignalArticleListResponseArticleJSON struct {
	ID              apijson.Field
	DatasetID       apijson.Field
	EventID         apijson.Field
	FeedDisplayName apijson.Field
	FeedID          apijson.Field
	FetchedAt       apijson.Field
	Link            apijson.Field
	PublishedAt     apijson.Field
	Read            apijson.Field
	ReadAt          apijson.Field
	Relevance       apijson.Field
	Severity        apijson.Field
	Summary         apijson.Field
	Tags            apijson.Field
	Title           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ThreatSignalArticleListResponseArticle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalArticleListResponseArticleJSON) RawJSON() string {
	return r.raw
}

// Beta. The article's relevance to the account's active Priority Intelligence
// Requirements. Null when the article has not been scored against them, including
// when the account has no active PIRs, or has not been scored against the
// requested pir_id.
type ThreatSignalArticleListResponseArticlesRelevance struct {
	// Beta. How relevant the article is to a Priority Intelligence Requirement: high,
	// medium, or low. Cloudflare may change the scores behind each band as relevance
	// scoring is calibrated.
	Band ThreatSignalArticleListResponseArticlesRelevanceBand `json:"band" api:"required"`
	// Beta. The pir_id requested from the article list, or otherwise the account's
	// most relevant active PIR for this article.
	PirID string `json:"pir_id" api:"required" format:"uuid"`
	// Beta. Every active PIR the article was scored against, most relevant first.
	// Draft and deleted PIRs are excluded.
	Pirs []ThreatSignalArticleListResponseArticlesRelevancePir `json:"pirs" api:"required"`
	JSON threatSignalArticleListResponseArticlesRelevanceJSON  `json:"-"`
}

// threatSignalArticleListResponseArticlesRelevanceJSON contains the JSON metadata
// for the struct [ThreatSignalArticleListResponseArticlesRelevance]
type threatSignalArticleListResponseArticlesRelevanceJSON struct {
	Band        apijson.Field
	PirID       apijson.Field
	Pirs        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalArticleListResponseArticlesRelevance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalArticleListResponseArticlesRelevanceJSON) RawJSON() string {
	return r.raw
}

// Beta. How relevant the article is to a Priority Intelligence Requirement: high,
// medium, or low. Cloudflare may change the scores behind each band as relevance
// scoring is calibrated.
type ThreatSignalArticleListResponseArticlesRelevanceBand string

const (
	ThreatSignalArticleListResponseArticlesRelevanceBandLow    ThreatSignalArticleListResponseArticlesRelevanceBand = "low"
	ThreatSignalArticleListResponseArticlesRelevanceBandMedium ThreatSignalArticleListResponseArticlesRelevanceBand = "medium"
	ThreatSignalArticleListResponseArticlesRelevanceBandHigh   ThreatSignalArticleListResponseArticlesRelevanceBand = "high"
)

func (r ThreatSignalArticleListResponseArticlesRelevanceBand) IsKnown() bool {
	switch r {
	case ThreatSignalArticleListResponseArticlesRelevanceBandLow, ThreatSignalArticleListResponseArticlesRelevanceBandMedium, ThreatSignalArticleListResponseArticlesRelevanceBandHigh:
		return true
	}
	return false
}

type ThreatSignalArticleListResponseArticlesRelevancePir struct {
	// Beta. How relevant the article is to a Priority Intelligence Requirement: high,
	// medium, or low. Cloudflare may change the scores behind each band as relevance
	// scoring is calibrated.
	Band  ThreatSignalArticleListResponseArticlesRelevancePirsBand `json:"band" api:"required"`
	PirID string                                                   `json:"pir_id" api:"required" format:"uuid"`
	JSON  threatSignalArticleListResponseArticlesRelevancePirJSON  `json:"-"`
}

// threatSignalArticleListResponseArticlesRelevancePirJSON contains the JSON
// metadata for the struct [ThreatSignalArticleListResponseArticlesRelevancePir]
type threatSignalArticleListResponseArticlesRelevancePirJSON struct {
	Band        apijson.Field
	PirID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalArticleListResponseArticlesRelevancePir) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalArticleListResponseArticlesRelevancePirJSON) RawJSON() string {
	return r.raw
}

// Beta. How relevant the article is to a Priority Intelligence Requirement: high,
// medium, or low. Cloudflare may change the scores behind each band as relevance
// scoring is calibrated.
type ThreatSignalArticleListResponseArticlesRelevancePirsBand string

const (
	ThreatSignalArticleListResponseArticlesRelevancePirsBandLow    ThreatSignalArticleListResponseArticlesRelevancePirsBand = "low"
	ThreatSignalArticleListResponseArticlesRelevancePirsBandMedium ThreatSignalArticleListResponseArticlesRelevancePirsBand = "medium"
	ThreatSignalArticleListResponseArticlesRelevancePirsBandHigh   ThreatSignalArticleListResponseArticlesRelevancePirsBand = "high"
)

func (r ThreatSignalArticleListResponseArticlesRelevancePirsBand) IsKnown() bool {
	switch r {
	case ThreatSignalArticleListResponseArticlesRelevancePirsBandLow, ThreatSignalArticleListResponseArticlesRelevancePirsBandMedium, ThreatSignalArticleListResponseArticlesRelevancePirsBandHigh:
		return true
	}
	return false
}

// Beta. Intrinsic article severity band, from informational to critical,
// independent of any organisation or Priority Intelligence Requirement. Null when
// the account has no active Priority Intelligence Requirements, scoring has not
// completed, or a historic article has not entered an automatic rescore window.
type ThreatSignalArticleListResponseArticlesSeverity string

const (
	ThreatSignalArticleListResponseArticlesSeverityInformational ThreatSignalArticleListResponseArticlesSeverity = "informational"
	ThreatSignalArticleListResponseArticlesSeverityLow           ThreatSignalArticleListResponseArticlesSeverity = "low"
	ThreatSignalArticleListResponseArticlesSeverityMedium        ThreatSignalArticleListResponseArticlesSeverity = "medium"
	ThreatSignalArticleListResponseArticlesSeverityHigh          ThreatSignalArticleListResponseArticlesSeverity = "high"
	ThreatSignalArticleListResponseArticlesSeverityCritical      ThreatSignalArticleListResponseArticlesSeverity = "critical"
)

func (r ThreatSignalArticleListResponseArticlesSeverity) IsKnown() bool {
	switch r {
	case ThreatSignalArticleListResponseArticlesSeverityInformational, ThreatSignalArticleListResponseArticlesSeverityLow, ThreatSignalArticleListResponseArticlesSeverityMedium, ThreatSignalArticleListResponseArticlesSeverityHigh, ThreatSignalArticleListResponseArticlesSeverityCritical:
		return true
	}
	return false
}

type ThreatSignalArticleListResponseArticlesTag struct {
	AppliedBy  ThreatSignalArticleListResponseArticlesTagsAppliedBy `json:"applied_by" api:"required"`
	CategoryID string                                               `json:"categoryId" api:"required,nullable" format:"uuid"`
	UUID       string                                               `json:"uuid" api:"required" format:"uuid"`
	Value      string                                               `json:"value" api:"required"`
	JSON       threatSignalArticleListResponseArticlesTagJSON       `json:"-"`
}

// threatSignalArticleListResponseArticlesTagJSON contains the JSON metadata for
// the struct [ThreatSignalArticleListResponseArticlesTag]
type threatSignalArticleListResponseArticlesTagJSON struct {
	AppliedBy   apijson.Field
	CategoryID  apijson.Field
	UUID        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalArticleListResponseArticlesTag) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalArticleListResponseArticlesTagJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalArticleListResponseArticlesTagsAppliedBy string

const (
	ThreatSignalArticleListResponseArticlesTagsAppliedByAI      ThreatSignalArticleListResponseArticlesTagsAppliedBy = "ai"
	ThreatSignalArticleListResponseArticlesTagsAppliedByAnalyst ThreatSignalArticleListResponseArticlesTagsAppliedBy = "analyst"
	ThreatSignalArticleListResponseArticlesTagsAppliedBySystem  ThreatSignalArticleListResponseArticlesTagsAppliedBy = "system"
)

func (r ThreatSignalArticleListResponseArticlesTagsAppliedBy) IsKnown() bool {
	switch r {
	case ThreatSignalArticleListResponseArticlesTagsAppliedByAI, ThreatSignalArticleListResponseArticlesTagsAppliedByAnalyst, ThreatSignalArticleListResponseArticlesTagsAppliedBySystem:
		return true
	}
	return false
}

type ThreatSignalArticleBulkEditResponse struct {
	UpdatedCount float64                                 `json:"updated_count" api:"required"`
	JSON         threatSignalArticleBulkEditResponseJSON `json:"-"`
}

// threatSignalArticleBulkEditResponseJSON contains the JSON metadata for the
// struct [ThreatSignalArticleBulkEditResponse]
type threatSignalArticleBulkEditResponseJSON struct {
	UpdatedCount apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ThreatSignalArticleBulkEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalArticleBulkEditResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalArticleEditResponse struct {
	ID              string                                      `json:"id" api:"required" format:"uuid"`
	BulletPoints    ThreatSignalArticleEditResponseBulletPoints `json:"bullet_points" api:"required,nullable"`
	ContentR2Key    string                                      `json:"content_r2_key" api:"required,nullable"`
	FeedDisplayName string                                      `json:"feed_display_name" api:"required,nullable"`
	FeedID          string                                      `json:"feed_id" api:"required" format:"uuid"`
	FetchedAt       string                                      `json:"fetched_at" api:"required"`
	// Progress of the article's indicator extraction and IOC contextualization run.
	// complete and failed are terminal; unknown means no run has been recorded.
	IndicatorExtractionStatus ThreatSignalArticleEditResponseIndicatorExtractionStatus `json:"indicator_extraction_status" api:"required"`
	Link                      string                                                   `json:"link" api:"required,nullable"`
	Metadata                  map[string]interface{}                                   `json:"metadata" api:"required,nullable"`
	PublishedAt               string                                                   `json:"published_at" api:"required,nullable"`
	Read                      bool                                                     `json:"read" api:"required"`
	ReadAt                    string                                                   `json:"read_at" api:"required,nullable"`
	// Beta. Intrinsic article severity band, from informational to critical,
	// independent of any organisation or Priority Intelligence Requirement. Null when
	// the account has no active Priority Intelligence Requirements, scoring has not
	// completed, or a historic article has not entered an automatic rescore window.
	Severity    ThreatSignalArticleEditResponseSeverity `json:"severity" api:"required,nullable"`
	SourceCount float64                                 `json:"source_count" api:"required"`
	// Persisted enrichment summary. Null until enrichment produces a summary.
	Summary         string                               `json:"summary" api:"required,nullable"`
	SummaryR2Key    string                               `json:"summary_r2_key" api:"required,nullable"`
	Tags            []ThreatSignalArticleEditResponseTag `json:"tags" api:"required"`
	Title           string                               `json:"title" api:"required,nullable"`
	SkillVersion    string                               `json:"skill_version" api:"nullable"`
	TagSkillVersion string                               `json:"tag_skill_version" api:"nullable"`
	JSON            threatSignalArticleEditResponseJSON  `json:"-"`
}

// threatSignalArticleEditResponseJSON contains the JSON metadata for the struct
// [ThreatSignalArticleEditResponse]
type threatSignalArticleEditResponseJSON struct {
	ID                        apijson.Field
	BulletPoints              apijson.Field
	ContentR2Key              apijson.Field
	FeedDisplayName           apijson.Field
	FeedID                    apijson.Field
	FetchedAt                 apijson.Field
	IndicatorExtractionStatus apijson.Field
	Link                      apijson.Field
	Metadata                  apijson.Field
	PublishedAt               apijson.Field
	Read                      apijson.Field
	ReadAt                    apijson.Field
	Severity                  apijson.Field
	SourceCount               apijson.Field
	Summary                   apijson.Field
	SummaryR2Key              apijson.Field
	Tags                      apijson.Field
	Title                     apijson.Field
	SkillVersion              apijson.Field
	TagSkillVersion           apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *ThreatSignalArticleEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalArticleEditResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalArticleEditResponseBulletPoints struct {
	Impact       string                                          `json:"impact" api:"required"`
	WhatHappened string                                          `json:"what_happened" api:"required"`
	WhoAffected  string                                          `json:"who_affected" api:"required"`
	JSON         threatSignalArticleEditResponseBulletPointsJSON `json:"-"`
}

// threatSignalArticleEditResponseBulletPointsJSON contains the JSON metadata for
// the struct [ThreatSignalArticleEditResponseBulletPoints]
type threatSignalArticleEditResponseBulletPointsJSON struct {
	Impact       apijson.Field
	WhatHappened apijson.Field
	WhoAffected  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ThreatSignalArticleEditResponseBulletPoints) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalArticleEditResponseBulletPointsJSON) RawJSON() string {
	return r.raw
}

// Progress of the article's indicator extraction and IOC contextualization run.
// complete and failed are terminal; unknown means no run has been recorded.
type ThreatSignalArticleEditResponseIndicatorExtractionStatus string

const (
	ThreatSignalArticleEditResponseIndicatorExtractionStatusInProgress ThreatSignalArticleEditResponseIndicatorExtractionStatus = "in_progress"
	ThreatSignalArticleEditResponseIndicatorExtractionStatusComplete   ThreatSignalArticleEditResponseIndicatorExtractionStatus = "complete"
	ThreatSignalArticleEditResponseIndicatorExtractionStatusFailed     ThreatSignalArticleEditResponseIndicatorExtractionStatus = "failed"
	ThreatSignalArticleEditResponseIndicatorExtractionStatusUnknown    ThreatSignalArticleEditResponseIndicatorExtractionStatus = "unknown"
)

func (r ThreatSignalArticleEditResponseIndicatorExtractionStatus) IsKnown() bool {
	switch r {
	case ThreatSignalArticleEditResponseIndicatorExtractionStatusInProgress, ThreatSignalArticleEditResponseIndicatorExtractionStatusComplete, ThreatSignalArticleEditResponseIndicatorExtractionStatusFailed, ThreatSignalArticleEditResponseIndicatorExtractionStatusUnknown:
		return true
	}
	return false
}

// Beta. Intrinsic article severity band, from informational to critical,
// independent of any organisation or Priority Intelligence Requirement. Null when
// the account has no active Priority Intelligence Requirements, scoring has not
// completed, or a historic article has not entered an automatic rescore window.
type ThreatSignalArticleEditResponseSeverity string

const (
	ThreatSignalArticleEditResponseSeverityInformational ThreatSignalArticleEditResponseSeverity = "informational"
	ThreatSignalArticleEditResponseSeverityLow           ThreatSignalArticleEditResponseSeverity = "low"
	ThreatSignalArticleEditResponseSeverityMedium        ThreatSignalArticleEditResponseSeverity = "medium"
	ThreatSignalArticleEditResponseSeverityHigh          ThreatSignalArticleEditResponseSeverity = "high"
	ThreatSignalArticleEditResponseSeverityCritical      ThreatSignalArticleEditResponseSeverity = "critical"
)

func (r ThreatSignalArticleEditResponseSeverity) IsKnown() bool {
	switch r {
	case ThreatSignalArticleEditResponseSeverityInformational, ThreatSignalArticleEditResponseSeverityLow, ThreatSignalArticleEditResponseSeverityMedium, ThreatSignalArticleEditResponseSeverityHigh, ThreatSignalArticleEditResponseSeverityCritical:
		return true
	}
	return false
}

type ThreatSignalArticleEditResponseTag struct {
	AppliedBy  ThreatSignalArticleEditResponseTagsAppliedBy `json:"applied_by" api:"required"`
	CategoryID string                                       `json:"categoryId" api:"required,nullable" format:"uuid"`
	UUID       string                                       `json:"uuid" api:"required" format:"uuid"`
	Value      string                                       `json:"value" api:"required"`
	JSON       threatSignalArticleEditResponseTagJSON       `json:"-"`
}

// threatSignalArticleEditResponseTagJSON contains the JSON metadata for the struct
// [ThreatSignalArticleEditResponseTag]
type threatSignalArticleEditResponseTagJSON struct {
	AppliedBy   apijson.Field
	CategoryID  apijson.Field
	UUID        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalArticleEditResponseTag) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalArticleEditResponseTagJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalArticleEditResponseTagsAppliedBy string

const (
	ThreatSignalArticleEditResponseTagsAppliedByAI      ThreatSignalArticleEditResponseTagsAppliedBy = "ai"
	ThreatSignalArticleEditResponseTagsAppliedByAnalyst ThreatSignalArticleEditResponseTagsAppliedBy = "analyst"
	ThreatSignalArticleEditResponseTagsAppliedBySystem  ThreatSignalArticleEditResponseTagsAppliedBy = "system"
)

func (r ThreatSignalArticleEditResponseTagsAppliedBy) IsKnown() bool {
	switch r {
	case ThreatSignalArticleEditResponseTagsAppliedByAI, ThreatSignalArticleEditResponseTagsAppliedByAnalyst, ThreatSignalArticleEditResponseTagsAppliedBySystem:
		return true
	}
	return false
}

type ThreatSignalArticleGetResponse struct {
	ID              string                                     `json:"id" api:"required" format:"uuid"`
	BulletPoints    ThreatSignalArticleGetResponseBulletPoints `json:"bullet_points" api:"required,nullable"`
	ContentR2Key    string                                     `json:"content_r2_key" api:"required,nullable"`
	FeedDisplayName string                                     `json:"feed_display_name" api:"required,nullable"`
	FeedID          string                                     `json:"feed_id" api:"required" format:"uuid"`
	FetchedAt       string                                     `json:"fetched_at" api:"required"`
	// Progress of the article's indicator extraction and IOC contextualization run.
	// complete and failed are terminal; unknown means no run has been recorded.
	IndicatorExtractionStatus ThreatSignalArticleGetResponseIndicatorExtractionStatus `json:"indicator_extraction_status" api:"required"`
	Link                      string                                                  `json:"link" api:"required,nullable"`
	Metadata                  map[string]interface{}                                  `json:"metadata" api:"required,nullable"`
	PublishedAt               string                                                  `json:"published_at" api:"required,nullable"`
	Read                      bool                                                    `json:"read" api:"required"`
	ReadAt                    string                                                  `json:"read_at" api:"required,nullable"`
	// Beta. The article's relevance to the account's active Priority Intelligence
	// Requirements. Null when the article has not been scored against them, including
	// when the account has no active PIRs, or has not been scored against the
	// requested pir_id.
	Relevance ThreatSignalArticleGetResponseRelevance `json:"relevance" api:"required,nullable"`
	// Beta. Intrinsic article severity band, from informational to critical,
	// independent of any organisation or Priority Intelligence Requirement. Null when
	// the account has no active Priority Intelligence Requirements, scoring has not
	// completed, or a historic article has not entered an automatic rescore window.
	Severity    ThreatSignalArticleGetResponseSeverity `json:"severity" api:"required,nullable"`
	SourceCount float64                                `json:"source_count" api:"required"`
	// Persisted enrichment summary. Null until enrichment produces a summary.
	Summary         string                              `json:"summary" api:"required,nullable"`
	SummaryR2Key    string                              `json:"summary_r2_key" api:"required,nullable"`
	Tags            []ThreatSignalArticleGetResponseTag `json:"tags" api:"required"`
	Title           string                              `json:"title" api:"required,nullable"`
	SkillVersion    string                              `json:"skill_version" api:"nullable"`
	TagSkillVersion string                              `json:"tag_skill_version" api:"nullable"`
	JSON            threatSignalArticleGetResponseJSON  `json:"-"`
}

// threatSignalArticleGetResponseJSON contains the JSON metadata for the struct
// [ThreatSignalArticleGetResponse]
type threatSignalArticleGetResponseJSON struct {
	ID                        apijson.Field
	BulletPoints              apijson.Field
	ContentR2Key              apijson.Field
	FeedDisplayName           apijson.Field
	FeedID                    apijson.Field
	FetchedAt                 apijson.Field
	IndicatorExtractionStatus apijson.Field
	Link                      apijson.Field
	Metadata                  apijson.Field
	PublishedAt               apijson.Field
	Read                      apijson.Field
	ReadAt                    apijson.Field
	Relevance                 apijson.Field
	Severity                  apijson.Field
	SourceCount               apijson.Field
	Summary                   apijson.Field
	SummaryR2Key              apijson.Field
	Tags                      apijson.Field
	Title                     apijson.Field
	SkillVersion              apijson.Field
	TagSkillVersion           apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *ThreatSignalArticleGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalArticleGetResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalArticleGetResponseBulletPoints struct {
	Impact       string                                         `json:"impact" api:"required"`
	WhatHappened string                                         `json:"what_happened" api:"required"`
	WhoAffected  string                                         `json:"who_affected" api:"required"`
	JSON         threatSignalArticleGetResponseBulletPointsJSON `json:"-"`
}

// threatSignalArticleGetResponseBulletPointsJSON contains the JSON metadata for
// the struct [ThreatSignalArticleGetResponseBulletPoints]
type threatSignalArticleGetResponseBulletPointsJSON struct {
	Impact       apijson.Field
	WhatHappened apijson.Field
	WhoAffected  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ThreatSignalArticleGetResponseBulletPoints) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalArticleGetResponseBulletPointsJSON) RawJSON() string {
	return r.raw
}

// Progress of the article's indicator extraction and IOC contextualization run.
// complete and failed are terminal; unknown means no run has been recorded.
type ThreatSignalArticleGetResponseIndicatorExtractionStatus string

const (
	ThreatSignalArticleGetResponseIndicatorExtractionStatusInProgress ThreatSignalArticleGetResponseIndicatorExtractionStatus = "in_progress"
	ThreatSignalArticleGetResponseIndicatorExtractionStatusComplete   ThreatSignalArticleGetResponseIndicatorExtractionStatus = "complete"
	ThreatSignalArticleGetResponseIndicatorExtractionStatusFailed     ThreatSignalArticleGetResponseIndicatorExtractionStatus = "failed"
	ThreatSignalArticleGetResponseIndicatorExtractionStatusUnknown    ThreatSignalArticleGetResponseIndicatorExtractionStatus = "unknown"
)

func (r ThreatSignalArticleGetResponseIndicatorExtractionStatus) IsKnown() bool {
	switch r {
	case ThreatSignalArticleGetResponseIndicatorExtractionStatusInProgress, ThreatSignalArticleGetResponseIndicatorExtractionStatusComplete, ThreatSignalArticleGetResponseIndicatorExtractionStatusFailed, ThreatSignalArticleGetResponseIndicatorExtractionStatusUnknown:
		return true
	}
	return false
}

// Beta. The article's relevance to the account's active Priority Intelligence
// Requirements. Null when the article has not been scored against them, including
// when the account has no active PIRs, or has not been scored against the
// requested pir_id.
type ThreatSignalArticleGetResponseRelevance struct {
	// Beta. How relevant the article is to a Priority Intelligence Requirement: high,
	// medium, or low. Cloudflare may change the scores behind each band as relevance
	// scoring is calibrated.
	Band ThreatSignalArticleGetResponseRelevanceBand `json:"band" api:"required"`
	// Beta. The pir_id requested from the article list, or otherwise the account's
	// most relevant active PIR for this article.
	PirID string `json:"pir_id" api:"required" format:"uuid"`
	// Beta. Every active PIR the article was scored against, most relevant first.
	// Draft and deleted PIRs are excluded.
	Pirs []ThreatSignalArticleGetResponseRelevancePir `json:"pirs" api:"required"`
	JSON threatSignalArticleGetResponseRelevanceJSON  `json:"-"`
}

// threatSignalArticleGetResponseRelevanceJSON contains the JSON metadata for the
// struct [ThreatSignalArticleGetResponseRelevance]
type threatSignalArticleGetResponseRelevanceJSON struct {
	Band        apijson.Field
	PirID       apijson.Field
	Pirs        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalArticleGetResponseRelevance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalArticleGetResponseRelevanceJSON) RawJSON() string {
	return r.raw
}

// Beta. How relevant the article is to a Priority Intelligence Requirement: high,
// medium, or low. Cloudflare may change the scores behind each band as relevance
// scoring is calibrated.
type ThreatSignalArticleGetResponseRelevanceBand string

const (
	ThreatSignalArticleGetResponseRelevanceBandLow    ThreatSignalArticleGetResponseRelevanceBand = "low"
	ThreatSignalArticleGetResponseRelevanceBandMedium ThreatSignalArticleGetResponseRelevanceBand = "medium"
	ThreatSignalArticleGetResponseRelevanceBandHigh   ThreatSignalArticleGetResponseRelevanceBand = "high"
)

func (r ThreatSignalArticleGetResponseRelevanceBand) IsKnown() bool {
	switch r {
	case ThreatSignalArticleGetResponseRelevanceBandLow, ThreatSignalArticleGetResponseRelevanceBandMedium, ThreatSignalArticleGetResponseRelevanceBandHigh:
		return true
	}
	return false
}

type ThreatSignalArticleGetResponseRelevancePir struct {
	// Beta. How relevant the article is to a Priority Intelligence Requirement: high,
	// medium, or low. Cloudflare may change the scores behind each band as relevance
	// scoring is calibrated.
	Band  ThreatSignalArticleGetResponseRelevancePirsBand `json:"band" api:"required"`
	PirID string                                          `json:"pir_id" api:"required" format:"uuid"`
	JSON  threatSignalArticleGetResponseRelevancePirJSON  `json:"-"`
}

// threatSignalArticleGetResponseRelevancePirJSON contains the JSON metadata for
// the struct [ThreatSignalArticleGetResponseRelevancePir]
type threatSignalArticleGetResponseRelevancePirJSON struct {
	Band        apijson.Field
	PirID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalArticleGetResponseRelevancePir) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalArticleGetResponseRelevancePirJSON) RawJSON() string {
	return r.raw
}

// Beta. How relevant the article is to a Priority Intelligence Requirement: high,
// medium, or low. Cloudflare may change the scores behind each band as relevance
// scoring is calibrated.
type ThreatSignalArticleGetResponseRelevancePirsBand string

const (
	ThreatSignalArticleGetResponseRelevancePirsBandLow    ThreatSignalArticleGetResponseRelevancePirsBand = "low"
	ThreatSignalArticleGetResponseRelevancePirsBandMedium ThreatSignalArticleGetResponseRelevancePirsBand = "medium"
	ThreatSignalArticleGetResponseRelevancePirsBandHigh   ThreatSignalArticleGetResponseRelevancePirsBand = "high"
)

func (r ThreatSignalArticleGetResponseRelevancePirsBand) IsKnown() bool {
	switch r {
	case ThreatSignalArticleGetResponseRelevancePirsBandLow, ThreatSignalArticleGetResponseRelevancePirsBandMedium, ThreatSignalArticleGetResponseRelevancePirsBandHigh:
		return true
	}
	return false
}

// Beta. Intrinsic article severity band, from informational to critical,
// independent of any organisation or Priority Intelligence Requirement. Null when
// the account has no active Priority Intelligence Requirements, scoring has not
// completed, or a historic article has not entered an automatic rescore window.
type ThreatSignalArticleGetResponseSeverity string

const (
	ThreatSignalArticleGetResponseSeverityInformational ThreatSignalArticleGetResponseSeverity = "informational"
	ThreatSignalArticleGetResponseSeverityLow           ThreatSignalArticleGetResponseSeverity = "low"
	ThreatSignalArticleGetResponseSeverityMedium        ThreatSignalArticleGetResponseSeverity = "medium"
	ThreatSignalArticleGetResponseSeverityHigh          ThreatSignalArticleGetResponseSeverity = "high"
	ThreatSignalArticleGetResponseSeverityCritical      ThreatSignalArticleGetResponseSeverity = "critical"
)

func (r ThreatSignalArticleGetResponseSeverity) IsKnown() bool {
	switch r {
	case ThreatSignalArticleGetResponseSeverityInformational, ThreatSignalArticleGetResponseSeverityLow, ThreatSignalArticleGetResponseSeverityMedium, ThreatSignalArticleGetResponseSeverityHigh, ThreatSignalArticleGetResponseSeverityCritical:
		return true
	}
	return false
}

type ThreatSignalArticleGetResponseTag struct {
	AppliedBy  ThreatSignalArticleGetResponseTagsAppliedBy `json:"applied_by" api:"required"`
	CategoryID string                                      `json:"categoryId" api:"required,nullable" format:"uuid"`
	UUID       string                                      `json:"uuid" api:"required" format:"uuid"`
	Value      string                                      `json:"value" api:"required"`
	JSON       threatSignalArticleGetResponseTagJSON       `json:"-"`
}

// threatSignalArticleGetResponseTagJSON contains the JSON metadata for the struct
// [ThreatSignalArticleGetResponseTag]
type threatSignalArticleGetResponseTagJSON struct {
	AppliedBy   apijson.Field
	CategoryID  apijson.Field
	UUID        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalArticleGetResponseTag) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalArticleGetResponseTagJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalArticleGetResponseTagsAppliedBy string

const (
	ThreatSignalArticleGetResponseTagsAppliedByAI      ThreatSignalArticleGetResponseTagsAppliedBy = "ai"
	ThreatSignalArticleGetResponseTagsAppliedByAnalyst ThreatSignalArticleGetResponseTagsAppliedBy = "analyst"
	ThreatSignalArticleGetResponseTagsAppliedBySystem  ThreatSignalArticleGetResponseTagsAppliedBy = "system"
)

func (r ThreatSignalArticleGetResponseTagsAppliedBy) IsKnown() bool {
	switch r {
	case ThreatSignalArticleGetResponseTagsAppliedByAI, ThreatSignalArticleGetResponseTagsAppliedByAnalyst, ThreatSignalArticleGetResponseTagsAppliedBySystem:
		return true
	}
	return false
}

type ThreatSignalArticleListParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Repeatable article UUID filter. Returns the union of matching account-owned
	// articles; use this to list every Threat Signals article referenced by an
	// indicator's sources.
	ArticleID param.Field[[]string] `query:"article_id" format:"uuid"`
	// Opaque cursor from a previous response's `next_cursor`. When provided,
	// pagination, ordering, totals, and article filters come from the cursor. Sending
	// `per_page`, `sort`, `include_total`, or any article filter alongside it returns
	// a 400 `CursorFilterConflictError`.
	Cursor param.Field[string] `query:"cursor"`
	// Exact feed-category filter after trimming. Omit to include every category.
	FeedCategory param.Field[string] `query:"feed_category"`
	// Feed UUID filter. Omit to include articles from every account feed.
	FeedID param.Field[string] `query:"feed_id" format:"uuid"`
	// Include articles fetched at or after this ISO 8601 date-time with Z or a numeric
	// UTC offset; normalized to UTC.
	FetchedAfter param.Field[time.Time] `query:"fetched_after" format:"date-time"`
	// Include articles fetched at or before this ISO 8601 date-time with Z or a
	// numeric UTC offset; normalized to UTC.
	FetchedBefore param.Field[time.Time] `query:"fetched_before" format:"date-time"`
	// When true on the first page, count matching articles and return an exact total.
	// Defaults to false; cursor pages do not recalculate totals.
	IncludeTotal param.Field[bool] `query:"include_total"`
	// Beta. Return only articles at or above this relevance band for pir_id, or for
	// any active PIR when pir_id is omitted. Excludes articles that have not been
	// scored.
	MinRelevance param.Field[ThreatSignalArticleListParamsMinRelevance] `query:"min_relevance"`
	// Number of articles per cursor page, from 1 to 100. Defaults to 20.
	PerPage param.Field[int64] `query:"per_page"`
	// Beta. Active Priority Intelligence Requirement to report, filter, and sort
	// relevance by. Omit to use each article's most relevant active PIR. Returns a 400
	// error unless the PIR is active.
	PirID param.Field[string] `query:"pir_id" format:"uuid"`
	// Include articles published at or after this ISO 8601 date-time with Z or a
	// numeric UTC offset; normalized to UTC.
	PublishedAfter param.Field[time.Time] `query:"published_after" format:"date-time"`
	// Include articles published at or before this ISO 8601 date-time with Z or a
	// numeric UTC offset; normalized to UTC.
	PublishedBefore param.Field[time.Time] `query:"published_before" format:"date-time"`
	// Filter by read state. Accepts true, false, 1, or 0; omit to include read and
	// unread articles.
	Read param.Field[bool] `query:"read"`
	// NFC-normalized, trimmed literal substring filter on article titles; ASCII
	// case-insensitive, from 1 to 500 characters.
	Search param.Field[string] `query:"search"`
	// Beta. Repeatable intrinsic severity band filter. An article matches any selected
	// band, the band reported as its severity. Excludes articles without a severity.
	// Like the severity field, severity is reported only while the account has an
	// active Priority Intelligence Requirement; otherwise no articles match.
	Severity param.Field[[]ThreatSignalArticleListParamsSeverity] `query:"severity"`
	// Sort by fetched_at, published_at, title, or relevance (beta). Prefix the field
	// with - for descending order; defaults to -fetched_at. -relevance lists the most
	// relevant articles first, by pir_id or by each article's most relevant active
	// PIR, with unscored articles last.
	Sort param.Field[string] `query:"sort"`
	// Filter by curated or custom feed source. Omit to include both source types.
	SourceType param.Field[ThreatSignalArticleListParamsSourceType] `query:"source_type"`
	// Legacy human-readable tag-value filter. Ignored when tag_id is supplied; prefer
	// tag_id.
	Tag param.Field[string] `query:"tag"`
	// Assignment provenance filter. When combined with tag_id or tag_category_id, the
	// matching assignment must have this provenance.
	TagAppliedBy param.Field[ThreatSignalArticleListParamsTagAppliedBy] `query:"tag_applied_by"`
	// Legacy category-name disambiguator for tag. It has no effect without tag; prefer
	// tag_category_id.
	TagCategory param.Field[string] `query:"tag_category"`
	// Repeatable tag-category UUID filter. An article matches any selected category;
	// when tag_id is also present, the tag and category groups are ANDed.
	TagCategoryID param.Field[[]string] `query:"tag_category_id" format:"uuid"`
	// Repeatable tag UUID filter. An article matches any selected tag.
	TagID param.Field[[]string] `query:"tag_id" format:"uuid"`
}

// URLQuery serializes [ThreatSignalArticleListParams]'s query parameters as
// `url.Values`.
func (r ThreatSignalArticleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Beta. Return only articles at or above this relevance band for pir_id, or for
// any active PIR when pir_id is omitted. Excludes articles that have not been
// scored.
type ThreatSignalArticleListParamsMinRelevance string

const (
	ThreatSignalArticleListParamsMinRelevanceLow    ThreatSignalArticleListParamsMinRelevance = "low"
	ThreatSignalArticleListParamsMinRelevanceMedium ThreatSignalArticleListParamsMinRelevance = "medium"
	ThreatSignalArticleListParamsMinRelevanceHigh   ThreatSignalArticleListParamsMinRelevance = "high"
)

func (r ThreatSignalArticleListParamsMinRelevance) IsKnown() bool {
	switch r {
	case ThreatSignalArticleListParamsMinRelevanceLow, ThreatSignalArticleListParamsMinRelevanceMedium, ThreatSignalArticleListParamsMinRelevanceHigh:
		return true
	}
	return false
}

type ThreatSignalArticleListParamsSeverity string

const (
	ThreatSignalArticleListParamsSeverityInformational ThreatSignalArticleListParamsSeverity = "informational"
	ThreatSignalArticleListParamsSeverityLow           ThreatSignalArticleListParamsSeverity = "low"
	ThreatSignalArticleListParamsSeverityMedium        ThreatSignalArticleListParamsSeverity = "medium"
	ThreatSignalArticleListParamsSeverityHigh          ThreatSignalArticleListParamsSeverity = "high"
	ThreatSignalArticleListParamsSeverityCritical      ThreatSignalArticleListParamsSeverity = "critical"
)

func (r ThreatSignalArticleListParamsSeverity) IsKnown() bool {
	switch r {
	case ThreatSignalArticleListParamsSeverityInformational, ThreatSignalArticleListParamsSeverityLow, ThreatSignalArticleListParamsSeverityMedium, ThreatSignalArticleListParamsSeverityHigh, ThreatSignalArticleListParamsSeverityCritical:
		return true
	}
	return false
}

// Filter by curated or custom feed source. Omit to include both source types.
type ThreatSignalArticleListParamsSourceType string

const (
	ThreatSignalArticleListParamsSourceTypeCurated ThreatSignalArticleListParamsSourceType = "curated"
	ThreatSignalArticleListParamsSourceTypeCustom  ThreatSignalArticleListParamsSourceType = "custom"
)

func (r ThreatSignalArticleListParamsSourceType) IsKnown() bool {
	switch r {
	case ThreatSignalArticleListParamsSourceTypeCurated, ThreatSignalArticleListParamsSourceTypeCustom:
		return true
	}
	return false
}

// Assignment provenance filter. When combined with tag_id or tag_category_id, the
// matching assignment must have this provenance.
type ThreatSignalArticleListParamsTagAppliedBy string

const (
	ThreatSignalArticleListParamsTagAppliedByAI      ThreatSignalArticleListParamsTagAppliedBy = "ai"
	ThreatSignalArticleListParamsTagAppliedByAnalyst ThreatSignalArticleListParamsTagAppliedBy = "analyst"
	ThreatSignalArticleListParamsTagAppliedBySystem  ThreatSignalArticleListParamsTagAppliedBy = "system"
)

func (r ThreatSignalArticleListParamsTagAppliedBy) IsKnown() bool {
	switch r {
	case ThreatSignalArticleListParamsTagAppliedByAI, ThreatSignalArticleListParamsTagAppliedByAnalyst, ThreatSignalArticleListParamsTagAppliedBySystem:
		return true
	}
	return false
}

type ThreatSignalArticleListResponseEnvelope struct {
	Errors  []ThreatSignalArticleListResponseEnvelopeErrors `json:"errors" api:"required"`
	Result  ThreatSignalArticleListResponse                 `json:"result" api:"required"`
	Success ThreatSignalArticleListResponseEnvelopeSuccess  `json:"success" api:"required"`
	JSON    threatSignalArticleListResponseEnvelopeJSON     `json:"-"`
}

// threatSignalArticleListResponseEnvelopeJSON contains the JSON metadata for the
// struct [ThreatSignalArticleListResponseEnvelope]
type threatSignalArticleListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalArticleListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalArticleListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalArticleListResponseEnvelopeErrors struct {
	Message string                                            `json:"message" api:"required"`
	JSON    threatSignalArticleListResponseEnvelopeErrorsJSON `json:"-"`
}

// threatSignalArticleListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ThreatSignalArticleListResponseEnvelopeErrors]
type threatSignalArticleListResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalArticleListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalArticleListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalArticleListResponseEnvelopeSuccess bool

const (
	ThreatSignalArticleListResponseEnvelopeSuccessTrue ThreatSignalArticleListResponseEnvelopeSuccess = true
)

func (r ThreatSignalArticleListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ThreatSignalArticleListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ThreatSignalArticleBulkEditParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Array of article UUIDs to mark (1-50).
	ArticleIDs param.Field[[]string] `json:"article_ids" api:"required" format:"uuid"`
	// True to mark read, false to mark unread.
	Read param.Field[bool] `json:"read" api:"required"`
}

func (r ThreatSignalArticleBulkEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ThreatSignalArticleBulkEditResponseEnvelope struct {
	Errors   []ThreatSignalArticleBulkEditResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ThreatSignalArticleBulkEditResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   ThreatSignalArticleBulkEditResponse                   `json:"result" api:"required"`
	Success  ThreatSignalArticleBulkEditResponseEnvelopeSuccess    `json:"success" api:"required"`
	JSON     threatSignalArticleBulkEditResponseEnvelopeJSON       `json:"-"`
}

// threatSignalArticleBulkEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [ThreatSignalArticleBulkEditResponseEnvelope]
type threatSignalArticleBulkEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalArticleBulkEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalArticleBulkEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalArticleBulkEditResponseEnvelopeErrors struct {
	Message string                                                `json:"message" api:"required"`
	JSON    threatSignalArticleBulkEditResponseEnvelopeErrorsJSON `json:"-"`
}

// threatSignalArticleBulkEditResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ThreatSignalArticleBulkEditResponseEnvelopeErrors]
type threatSignalArticleBulkEditResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalArticleBulkEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalArticleBulkEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalArticleBulkEditResponseEnvelopeMessages struct {
	Message string  `json:"message" api:"required"`
	Code    float64 `json:"code"`
	// The current count of custom feeds for the account.
	CustomFeedCount int64 `json:"custom_feed_count"`
	// The custom feed limit for the account.
	CustomFeedLimit int64                                                     `json:"custom_feed_limit"`
	Expected        string                                                    `json:"expected"`
	Path            []string                                                  `json:"path"`
	Reason          ThreatSignalArticleBulkEditResponseEnvelopeMessagesReason `json:"reason"`
	JSON            threatSignalArticleBulkEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// threatSignalArticleBulkEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ThreatSignalArticleBulkEditResponseEnvelopeMessages]
type threatSignalArticleBulkEditResponseEnvelopeMessagesJSON struct {
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

func (r *ThreatSignalArticleBulkEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalArticleBulkEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalArticleBulkEditResponseEnvelopeMessagesReason string

const (
	ThreatSignalArticleBulkEditResponseEnvelopeMessagesReasonThreatSignalsFeedLimit           ThreatSignalArticleBulkEditResponseEnvelopeMessagesReason = "threat_signals_feed_limit"
	ThreatSignalArticleBulkEditResponseEnvelopeMessagesReasonFreeCustomSkillsDisabled         ThreatSignalArticleBulkEditResponseEnvelopeMessagesReason = "free_custom_skills_disabled"
	ThreatSignalArticleBulkEditResponseEnvelopeMessagesReasonFreeSkillRunDisabled             ThreatSignalArticleBulkEditResponseEnvelopeMessagesReason = "free_skill_run_disabled"
	ThreatSignalArticleBulkEditResponseEnvelopeMessagesReasonFreeTierReconciliationInProgress ThreatSignalArticleBulkEditResponseEnvelopeMessagesReason = "free_tier_reconciliation_in_progress"
	ThreatSignalArticleBulkEditResponseEnvelopeMessagesReasonFreeTierReconciliationFailed     ThreatSignalArticleBulkEditResponseEnvelopeMessagesReason = "free_tier_reconciliation_failed"
	ThreatSignalArticleBulkEditResponseEnvelopeMessagesReasonRawContentExpired                ThreatSignalArticleBulkEditResponseEnvelopeMessagesReason = "raw_content_expired"
	ThreatSignalArticleBulkEditResponseEnvelopeMessagesReasonCuratedVisibilityChanged         ThreatSignalArticleBulkEditResponseEnvelopeMessagesReason = "curated_visibility_changed"
)

func (r ThreatSignalArticleBulkEditResponseEnvelopeMessagesReason) IsKnown() bool {
	switch r {
	case ThreatSignalArticleBulkEditResponseEnvelopeMessagesReasonThreatSignalsFeedLimit, ThreatSignalArticleBulkEditResponseEnvelopeMessagesReasonFreeCustomSkillsDisabled, ThreatSignalArticleBulkEditResponseEnvelopeMessagesReasonFreeSkillRunDisabled, ThreatSignalArticleBulkEditResponseEnvelopeMessagesReasonFreeTierReconciliationInProgress, ThreatSignalArticleBulkEditResponseEnvelopeMessagesReasonFreeTierReconciliationFailed, ThreatSignalArticleBulkEditResponseEnvelopeMessagesReasonRawContentExpired, ThreatSignalArticleBulkEditResponseEnvelopeMessagesReasonCuratedVisibilityChanged:
		return true
	}
	return false
}

type ThreatSignalArticleBulkEditResponseEnvelopeSuccess bool

const (
	ThreatSignalArticleBulkEditResponseEnvelopeSuccessTrue ThreatSignalArticleBulkEditResponseEnvelopeSuccess = true
)

func (r ThreatSignalArticleBulkEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ThreatSignalArticleBulkEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ThreatSignalArticleEditParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// True to mark the article read, false to mark unread.
	Read param.Field[bool] `json:"read" api:"required"`
}

func (r ThreatSignalArticleEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ThreatSignalArticleEditResponseEnvelope struct {
	Errors  []ThreatSignalArticleEditResponseEnvelopeErrors `json:"errors" api:"required"`
	Result  ThreatSignalArticleEditResponse                 `json:"result" api:"required"`
	Success ThreatSignalArticleEditResponseEnvelopeSuccess  `json:"success" api:"required"`
	JSON    threatSignalArticleEditResponseEnvelopeJSON     `json:"-"`
}

// threatSignalArticleEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [ThreatSignalArticleEditResponseEnvelope]
type threatSignalArticleEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalArticleEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalArticleEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalArticleEditResponseEnvelopeErrors struct {
	Message string                                            `json:"message" api:"required"`
	JSON    threatSignalArticleEditResponseEnvelopeErrorsJSON `json:"-"`
}

// threatSignalArticleEditResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ThreatSignalArticleEditResponseEnvelopeErrors]
type threatSignalArticleEditResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalArticleEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalArticleEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalArticleEditResponseEnvelopeSuccess bool

const (
	ThreatSignalArticleEditResponseEnvelopeSuccessTrue ThreatSignalArticleEditResponseEnvelopeSuccess = true
)

func (r ThreatSignalArticleEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ThreatSignalArticleEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ThreatSignalArticleGetParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type ThreatSignalArticleGetResponseEnvelope struct {
	Errors  []ThreatSignalArticleGetResponseEnvelopeErrors `json:"errors" api:"required"`
	Result  ThreatSignalArticleGetResponse                 `json:"result" api:"required"`
	Success ThreatSignalArticleGetResponseEnvelopeSuccess  `json:"success" api:"required"`
	JSON    threatSignalArticleGetResponseEnvelopeJSON     `json:"-"`
}

// threatSignalArticleGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ThreatSignalArticleGetResponseEnvelope]
type threatSignalArticleGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalArticleGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalArticleGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalArticleGetResponseEnvelopeErrors struct {
	Message string                                           `json:"message" api:"required"`
	JSON    threatSignalArticleGetResponseEnvelopeErrorsJSON `json:"-"`
}

// threatSignalArticleGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ThreatSignalArticleGetResponseEnvelopeErrors]
type threatSignalArticleGetResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalArticleGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalArticleGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalArticleGetResponseEnvelopeSuccess bool

const (
	ThreatSignalArticleGetResponseEnvelopeSuccessTrue ThreatSignalArticleGetResponseEnvelopeSuccess = true
)

func (r ThreatSignalArticleGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ThreatSignalArticleGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
