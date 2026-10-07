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

// ThreatSignalFeedService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreatSignalFeedService] method instead.
type ThreatSignalFeedService struct {
	Options []option.RequestOption
	Raw     *ThreatSignalFeedRawService
	Skills  *ThreatSignalFeedSkillService
}

// NewThreatSignalFeedService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewThreatSignalFeedService(opts ...option.RequestOption) (r *ThreatSignalFeedService) {
	r = &ThreatSignalFeedService{}
	r.Options = opts
	r.Raw = NewThreatSignalFeedRawService(opts...)
	r.Skills = NewThreatSignalFeedSkillService(opts...)
	return
}

// Subscribes the account to a custom or curated Threat Signals feed.
func (r *ThreatSignalFeedService) New(ctx context.Context, params ThreatSignalFeedNewParams, opts ...option.RequestOption) (res *ThreatSignalFeedNewResponse, err error) {
	var env ThreatSignalFeedNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/v2/threat-signals/feeds", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Lists the account's Threat Signals feed subscriptions.
func (r *ThreatSignalFeedService) List(ctx context.Context, params ThreatSignalFeedListParams, opts ...option.RequestOption) (res *pagination.V4PagePagination[ThreatSignalFeedListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/v2/threat-signals/feeds", params.AccountID)
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

// Lists the account's Threat Signals feed subscriptions.
func (r *ThreatSignalFeedService) ListAutoPaging(ctx context.Context, params ThreatSignalFeedListParams, opts ...option.RequestOption) *pagination.V4PagePaginationAutoPager[ThreatSignalFeedListResponse] {
	return pagination.NewV4PagePaginationAutoPager(r.List(ctx, params, opts...))
}

// Unsubscribes the account from a Threat Signals feed and deletes its articles.
func (r *ThreatSignalFeedService) Delete(ctx context.Context, feedID string, body ThreatSignalFeedDeleteParams, opts ...option.RequestOption) (res *ThreatSignalFeedDeleteResponse, err error) {
	var env ThreatSignalFeedDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if feedID == "" {
		err = errors.New("missing required feed_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/v2/threat-signals/feeds/%s", body.AccountID, feedID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Updates a Threat Signals feed subscription.
func (r *ThreatSignalFeedService) Edit(ctx context.Context, feedID string, params ThreatSignalFeedEditParams, opts ...option.RequestOption) (res *ThreatSignalFeedEditResponse, err error) {
	var env ThreatSignalFeedEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if feedID == "" {
		err = errors.New("missing required feed_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/v2/threat-signals/feeds/%s", params.AccountID, feedID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Starts an immediate poll of one or all Threat Signals feeds.
func (r *ThreatSignalFeedService) Poll(ctx context.Context, params ThreatSignalFeedPollParams, opts ...option.RequestOption) (res *ThreatSignalFeedPollResponse, err error) {
	var env ThreatSignalFeedPollResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/v2/threat-signals/feeds/poll", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type ThreatSignalFeedNewResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Feed category identifier. Null when unset.
	CategoryID string `json:"category_id" api:"required,nullable"`
	// Display name of the feed category. Null when unset or unresolvable.
	CategoryName string `json:"category_name" api:"required,nullable"`
	CreatedAt    string `json:"created_at" api:"required"`
	// Curated catalog feed this subscription was created from. Null for custom feeds.
	CuratedFeedID string `json:"curated_feed_id" api:"required,nullable"`
	DisplayName   string `json:"display_name" api:"required,nullable"`
	Enabled       bool   `json:"enabled" api:"required"`
	LastPolledAt  string `json:"last_polled_at" api:"required,nullable"`
	PollIntervalS int64  `json:"poll_interval_s" api:"required"`
	// `custom` for a feed added by URL, `curated` for a curated catalog feed.
	SourceType string `json:"source_type" api:"required"`
	// Polling health: `active`, or `error` after a failed poll.
	Status       string                          `json:"status" api:"required"`
	SubscribedAt string                          `json:"subscribed_at" api:"required,nullable"`
	Title        string                          `json:"title" api:"required,nullable"`
	UpdatedAt    string                          `json:"updated_at" api:"required"`
	URL          string                          `json:"url" api:"required"`
	JSON         threatSignalFeedNewResponseJSON `json:"-"`
}

// threatSignalFeedNewResponseJSON contains the JSON metadata for the struct
// [ThreatSignalFeedNewResponse]
type threatSignalFeedNewResponseJSON struct {
	ID            apijson.Field
	CategoryID    apijson.Field
	CategoryName  apijson.Field
	CreatedAt     apijson.Field
	CuratedFeedID apijson.Field
	DisplayName   apijson.Field
	Enabled       apijson.Field
	LastPolledAt  apijson.Field
	PollIntervalS apijson.Field
	SourceType    apijson.Field
	Status        apijson.Field
	SubscribedAt  apijson.Field
	Title         apijson.Field
	UpdatedAt     apijson.Field
	URL           apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ThreatSignalFeedNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalFeedNewResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalFeedListResponse struct {
	// Number of feeds on this page.
	Count      int64                              `json:"count" api:"required"`
	Feeds      []ThreatSignalFeedListResponseFeed `json:"feeds" api:"required"`
	Page       int64                              `json:"page" api:"required"`
	PerPage    int64                              `json:"per_page" api:"required"`
	TotalCount int64                              `json:"total_count" api:"required"`
	JSON       threatSignalFeedListResponseJSON   `json:"-"`
}

// threatSignalFeedListResponseJSON contains the JSON metadata for the struct
// [ThreatSignalFeedListResponse]
type threatSignalFeedListResponseJSON struct {
	Count       apijson.Field
	Feeds       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalFeedListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalFeedListResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalFeedListResponseFeed struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Feed category identifier. Null when unset.
	CategoryID string `json:"category_id" api:"required,nullable"`
	// Display name of the feed category. Null when unset or unresolvable.
	CategoryName string `json:"category_name" api:"required,nullable"`
	CreatedAt    string `json:"created_at" api:"required"`
	// Curated catalog feed this subscription was created from. Null for custom feeds.
	CuratedFeedID string `json:"curated_feed_id" api:"required,nullable"`
	DisplayName   string `json:"display_name" api:"required,nullable"`
	Enabled       bool   `json:"enabled" api:"required"`
	LastPolledAt  string `json:"last_polled_at" api:"required,nullable"`
	PollIntervalS int64  `json:"poll_interval_s" api:"required"`
	// `custom` for a feed added by URL, `curated` for a curated catalog feed.
	SourceType string `json:"source_type" api:"required"`
	// Polling health: `active`, or `error` after a failed poll.
	Status       string                               `json:"status" api:"required"`
	SubscribedAt string                               `json:"subscribed_at" api:"required,nullable"`
	Title        string                               `json:"title" api:"required,nullable"`
	UpdatedAt    string                               `json:"updated_at" api:"required"`
	URL          string                               `json:"url" api:"required"`
	JSON         threatSignalFeedListResponseFeedJSON `json:"-"`
}

// threatSignalFeedListResponseFeedJSON contains the JSON metadata for the struct
// [ThreatSignalFeedListResponseFeed]
type threatSignalFeedListResponseFeedJSON struct {
	ID            apijson.Field
	CategoryID    apijson.Field
	CategoryName  apijson.Field
	CreatedAt     apijson.Field
	CuratedFeedID apijson.Field
	DisplayName   apijson.Field
	Enabled       apijson.Field
	LastPolledAt  apijson.Field
	PollIntervalS apijson.Field
	SourceType    apijson.Field
	Status        apijson.Field
	SubscribedAt  apijson.Field
	Title         apijson.Field
	UpdatedAt     apijson.Field
	URL           apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ThreatSignalFeedListResponseFeed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalFeedListResponseFeedJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalFeedDeleteResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Feed category identifier. Null when unset.
	CategoryID string `json:"category_id" api:"required,nullable"`
	// Display name of the feed category. Null when unset or unresolvable.
	CategoryName string `json:"category_name" api:"required,nullable"`
	CreatedAt    string `json:"created_at" api:"required"`
	// Curated catalog feed this subscription was created from. Null for custom feeds.
	CuratedFeedID string `json:"curated_feed_id" api:"required,nullable"`
	DisplayName   string `json:"display_name" api:"required,nullable"`
	Enabled       bool   `json:"enabled" api:"required"`
	LastPolledAt  string `json:"last_polled_at" api:"required,nullable"`
	PollIntervalS int64  `json:"poll_interval_s" api:"required"`
	// `custom` for a feed added by URL, `curated` for a curated catalog feed.
	SourceType string `json:"source_type" api:"required"`
	// Polling health: `active`, or `error` after a failed poll.
	Status       string                             `json:"status" api:"required"`
	SubscribedAt string                             `json:"subscribed_at" api:"required,nullable"`
	Title        string                             `json:"title" api:"required,nullable"`
	UpdatedAt    string                             `json:"updated_at" api:"required"`
	URL          string                             `json:"url" api:"required"`
	JSON         threatSignalFeedDeleteResponseJSON `json:"-"`
}

// threatSignalFeedDeleteResponseJSON contains the JSON metadata for the struct
// [ThreatSignalFeedDeleteResponse]
type threatSignalFeedDeleteResponseJSON struct {
	ID            apijson.Field
	CategoryID    apijson.Field
	CategoryName  apijson.Field
	CreatedAt     apijson.Field
	CuratedFeedID apijson.Field
	DisplayName   apijson.Field
	Enabled       apijson.Field
	LastPolledAt  apijson.Field
	PollIntervalS apijson.Field
	SourceType    apijson.Field
	Status        apijson.Field
	SubscribedAt  apijson.Field
	Title         apijson.Field
	UpdatedAt     apijson.Field
	URL           apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ThreatSignalFeedDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalFeedDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalFeedEditResponse struct {
	ID string `json:"id" api:"required" format:"uuid"`
	// Feed category identifier. Null when unset.
	CategoryID string `json:"category_id" api:"required,nullable"`
	// Display name of the feed category. Null when unset or unresolvable.
	CategoryName string `json:"category_name" api:"required,nullable"`
	CreatedAt    string `json:"created_at" api:"required"`
	// Curated catalog feed this subscription was created from. Null for custom feeds.
	CuratedFeedID string `json:"curated_feed_id" api:"required,nullable"`
	DisplayName   string `json:"display_name" api:"required,nullable"`
	Enabled       bool   `json:"enabled" api:"required"`
	LastPolledAt  string `json:"last_polled_at" api:"required,nullable"`
	PollIntervalS int64  `json:"poll_interval_s" api:"required"`
	// `custom` for a feed added by URL, `curated` for a curated catalog feed.
	SourceType string `json:"source_type" api:"required"`
	// Polling health: `active`, or `error` after a failed poll.
	Status       string                           `json:"status" api:"required"`
	SubscribedAt string                           `json:"subscribed_at" api:"required,nullable"`
	Title        string                           `json:"title" api:"required,nullable"`
	UpdatedAt    string                           `json:"updated_at" api:"required"`
	URL          string                           `json:"url" api:"required"`
	JSON         threatSignalFeedEditResponseJSON `json:"-"`
}

// threatSignalFeedEditResponseJSON contains the JSON metadata for the struct
// [ThreatSignalFeedEditResponse]
type threatSignalFeedEditResponseJSON struct {
	ID            apijson.Field
	CategoryID    apijson.Field
	CategoryName  apijson.Field
	CreatedAt     apijson.Field
	CuratedFeedID apijson.Field
	DisplayName   apijson.Field
	Enabled       apijson.Field
	LastPolledAt  apijson.Field
	PollIntervalS apijson.Field
	SourceType    apijson.Field
	Status        apijson.Field
	SubscribedAt  apijson.Field
	Title         apijson.Field
	UpdatedAt     apijson.Field
	URL           apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ThreatSignalFeedEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalFeedEditResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalFeedPollResponse struct {
	Errors    float64                            `json:"errors" api:"required"`
	Feeds     []ThreatSignalFeedPollResponseFeed `json:"feeds" api:"required"`
	Triggered float64                            `json:"triggered" api:"required"`
	JSON      threatSignalFeedPollResponseJSON   `json:"-"`
}

// threatSignalFeedPollResponseJSON contains the JSON metadata for the struct
// [ThreatSignalFeedPollResponse]
type threatSignalFeedPollResponseJSON struct {
	Errors      apijson.Field
	Feeds       apijson.Field
	Triggered   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalFeedPollResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalFeedPollResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalFeedPollResponseFeed struct {
	FeedID      string                                  `json:"feed_id" api:"required" format:"uuid"`
	Status      ThreatSignalFeedPollResponseFeedsStatus `json:"status" api:"required"`
	WorkflowID  string                                  `json:"workflow_id" api:"required"`
	FeedEnabled bool                                    `json:"feed_enabled"`
	JSON        threatSignalFeedPollResponseFeedJSON    `json:"-"`
}

// threatSignalFeedPollResponseFeedJSON contains the JSON metadata for the struct
// [ThreatSignalFeedPollResponseFeed]
type threatSignalFeedPollResponseFeedJSON struct {
	FeedID      apijson.Field
	Status      apijson.Field
	WorkflowID  apijson.Field
	FeedEnabled apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalFeedPollResponseFeed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalFeedPollResponseFeedJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalFeedPollResponseFeedsStatus string

const (
	ThreatSignalFeedPollResponseFeedsStatusWorkflowCreated ThreatSignalFeedPollResponseFeedsStatus = "workflow_created"
	ThreatSignalFeedPollResponseFeedsStatusError           ThreatSignalFeedPollResponseFeedsStatus = "error"
)

func (r ThreatSignalFeedPollResponseFeedsStatus) IsKnown() bool {
	switch r {
	case ThreatSignalFeedPollResponseFeedsStatusWorkflowCreated, ThreatSignalFeedPollResponseFeedsStatusError:
		return true
	}
	return false
}

type ThreatSignalFeedNewParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// One of the predefined Threat Signals feed categories; see GET
	// /:account_id/v2/threat-signals/categories.
	CategoryID    param.Field[ThreatSignalFeedNewParamsCategoryID] `json:"category_id"`
	CuratedFeedID param.Field[string]                              `json:"curated_feed_id" format:"uuid"`
	DisplayName   param.Field[string]                              `json:"display_name"`
	Enabled       param.Field[bool]                                `json:"enabled"`
	PollIntervalS param.Field[int64]                               `json:"poll_interval_s"`
	Title         param.Field[string]                              `json:"title"`
	URL           param.Field[string]                              `json:"url" format:"uri"`
}

func (r ThreatSignalFeedNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// One of the predefined Threat Signals feed categories; see GET
// /:account_id/v2/threat-signals/categories.
type ThreatSignalFeedNewParamsCategoryID string

const (
	ThreatSignalFeedNewParamsCategoryIDB12a0fd6F7b9_5393_9ef3F888d506c550  ThreatSignalFeedNewParamsCategoryID = "b12a0fd6-f7b9-5393-9ef3-f888d506c550"
	ThreatSignalFeedNewParamsCategoryIDD5b70eaa626f5761B55b6d9590df49fb    ThreatSignalFeedNewParamsCategoryID = "d5b70eaa-626f-5761-b55b-6d9590df49fb"
	ThreatSignalFeedNewParamsCategoryID3b572d2b890d5286_9433F18c85079030   ThreatSignalFeedNewParamsCategoryID = "3b572d2b-890d-5286-9433-f18c85079030"
	ThreatSignalFeedNewParamsCategoryID17f90d3b37d3_5241_8ad4_7d6abbc2006c ThreatSignalFeedNewParamsCategoryID = "17f90d3b-37d3-5241-8ad4-7d6abbc2006c"
	ThreatSignalFeedNewParamsCategoryIDC68f28e9_7e8f5d4b853bF3076893a9ee   ThreatSignalFeedNewParamsCategoryID = "c68f28e9-7e8f-5d4b-853b-f3076893a9ee"
	ThreatSignalFeedNewParamsCategoryIDBb0e4a94_38ab5c14_80a7_28cee9f4b139 ThreatSignalFeedNewParamsCategoryID = "bb0e4a94-38ab-5c14-80a7-28cee9f4b139"
	ThreatSignalFeedNewParamsCategoryIDB1ef66d9A73c58dcB269_22d34dfd11f4   ThreatSignalFeedNewParamsCategoryID = "b1ef66d9-a73c-58dc-b269-22d34dfd11f4"
	ThreatSignalFeedNewParamsCategoryIDAb02a976_0a20_5c76A553_7f6325afacfe ThreatSignalFeedNewParamsCategoryID = "ab02a976-0a20-5c76-a553-7f6325afacfe"
)

func (r ThreatSignalFeedNewParamsCategoryID) IsKnown() bool {
	switch r {
	case ThreatSignalFeedNewParamsCategoryIDB12a0fd6F7b9_5393_9ef3F888d506c550, ThreatSignalFeedNewParamsCategoryIDD5b70eaa626f5761B55b6d9590df49fb, ThreatSignalFeedNewParamsCategoryID3b572d2b890d5286_9433F18c85079030, ThreatSignalFeedNewParamsCategoryID17f90d3b37d3_5241_8ad4_7d6abbc2006c, ThreatSignalFeedNewParamsCategoryIDC68f28e9_7e8f5d4b853bF3076893a9ee, ThreatSignalFeedNewParamsCategoryIDBb0e4a94_38ab5c14_80a7_28cee9f4b139, ThreatSignalFeedNewParamsCategoryIDB1ef66d9A73c58dcB269_22d34dfd11f4, ThreatSignalFeedNewParamsCategoryIDAb02a976_0a20_5c76A553_7f6325afacfe:
		return true
	}
	return false
}

type ThreatSignalFeedNewResponseEnvelope struct {
	Errors   []ThreatSignalFeedNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ThreatSignalFeedNewResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   ThreatSignalFeedNewResponse                   `json:"result" api:"required"`
	Success  ThreatSignalFeedNewResponseEnvelopeSuccess    `json:"success" api:"required"`
	JSON     threatSignalFeedNewResponseEnvelopeJSON       `json:"-"`
}

// threatSignalFeedNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [ThreatSignalFeedNewResponseEnvelope]
type threatSignalFeedNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalFeedNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalFeedNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalFeedNewResponseEnvelopeErrors struct {
	Message string                                        `json:"message" api:"required"`
	JSON    threatSignalFeedNewResponseEnvelopeErrorsJSON `json:"-"`
}

// threatSignalFeedNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ThreatSignalFeedNewResponseEnvelopeErrors]
type threatSignalFeedNewResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalFeedNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalFeedNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalFeedNewResponseEnvelopeMessages struct {
	Message  string                                            `json:"message" api:"required"`
	Code     float64                                           `json:"code"`
	Expected string                                            `json:"expected"`
	Path     []string                                          `json:"path"`
	Reason   ThreatSignalFeedNewResponseEnvelopeMessagesReason `json:"reason"`
	JSON     threatSignalFeedNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// threatSignalFeedNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ThreatSignalFeedNewResponseEnvelopeMessages]
type threatSignalFeedNewResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Expected    apijson.Field
	Path        apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalFeedNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalFeedNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalFeedNewResponseEnvelopeMessagesReason string

const (
	ThreatSignalFeedNewResponseEnvelopeMessagesReasonFreeCustomFeedLimit              ThreatSignalFeedNewResponseEnvelopeMessagesReason = "free_custom_feed_limit"
	ThreatSignalFeedNewResponseEnvelopeMessagesReasonFreeCustomSkillsDisabled         ThreatSignalFeedNewResponseEnvelopeMessagesReason = "free_custom_skills_disabled"
	ThreatSignalFeedNewResponseEnvelopeMessagesReasonFreeTierReconciliationInProgress ThreatSignalFeedNewResponseEnvelopeMessagesReason = "free_tier_reconciliation_in_progress"
	ThreatSignalFeedNewResponseEnvelopeMessagesReasonFreeTierReconciliationFailed     ThreatSignalFeedNewResponseEnvelopeMessagesReason = "free_tier_reconciliation_failed"
	ThreatSignalFeedNewResponseEnvelopeMessagesReasonRawContentExpired                ThreatSignalFeedNewResponseEnvelopeMessagesReason = "raw_content_expired"
	ThreatSignalFeedNewResponseEnvelopeMessagesReasonCuratedVisibilityChanged         ThreatSignalFeedNewResponseEnvelopeMessagesReason = "curated_visibility_changed"
)

func (r ThreatSignalFeedNewResponseEnvelopeMessagesReason) IsKnown() bool {
	switch r {
	case ThreatSignalFeedNewResponseEnvelopeMessagesReasonFreeCustomFeedLimit, ThreatSignalFeedNewResponseEnvelopeMessagesReasonFreeCustomSkillsDisabled, ThreatSignalFeedNewResponseEnvelopeMessagesReasonFreeTierReconciliationInProgress, ThreatSignalFeedNewResponseEnvelopeMessagesReasonFreeTierReconciliationFailed, ThreatSignalFeedNewResponseEnvelopeMessagesReasonRawContentExpired, ThreatSignalFeedNewResponseEnvelopeMessagesReasonCuratedVisibilityChanged:
		return true
	}
	return false
}

type ThreatSignalFeedNewResponseEnvelopeSuccess bool

const (
	ThreatSignalFeedNewResponseEnvelopeSuccessTrue ThreatSignalFeedNewResponseEnvelopeSuccess = true
)

func (r ThreatSignalFeedNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ThreatSignalFeedNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ThreatSignalFeedListParams struct {
	AccountID  param.Field[string]                               `path:"account_id" api:"required"`
	Category   param.Field[string]                               `query:"category"`
	Enabled    param.Field[bool]                                 `query:"enabled"`
	Limit      param.Field[int64]                                `query:"limit"`
	Page       param.Field[int64]                                `query:"page"`
	PerPage    param.Field[int64]                                `query:"per_page"`
	Sort       param.Field[string]                               `query:"sort"`
	SourceType param.Field[ThreatSignalFeedListParamsSourceType] `query:"source_type"`
	Status     param.Field[string]                               `query:"status"`
}

// URLQuery serializes [ThreatSignalFeedListParams]'s query parameters as
// `url.Values`.
func (r ThreatSignalFeedListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ThreatSignalFeedListParamsSourceType string

const (
	ThreatSignalFeedListParamsSourceTypeCurated ThreatSignalFeedListParamsSourceType = "curated"
	ThreatSignalFeedListParamsSourceTypeCustom  ThreatSignalFeedListParamsSourceType = "custom"
)

func (r ThreatSignalFeedListParamsSourceType) IsKnown() bool {
	switch r {
	case ThreatSignalFeedListParamsSourceTypeCurated, ThreatSignalFeedListParamsSourceTypeCustom:
		return true
	}
	return false
}

type ThreatSignalFeedDeleteParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type ThreatSignalFeedDeleteResponseEnvelope struct {
	Errors   []ThreatSignalFeedDeleteResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ThreatSignalFeedDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   ThreatSignalFeedDeleteResponse                   `json:"result" api:"required"`
	Success  ThreatSignalFeedDeleteResponseEnvelopeSuccess    `json:"success" api:"required"`
	JSON     threatSignalFeedDeleteResponseEnvelopeJSON       `json:"-"`
}

// threatSignalFeedDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [ThreatSignalFeedDeleteResponseEnvelope]
type threatSignalFeedDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalFeedDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalFeedDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalFeedDeleteResponseEnvelopeErrors struct {
	Message string                                           `json:"message" api:"required"`
	JSON    threatSignalFeedDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// threatSignalFeedDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ThreatSignalFeedDeleteResponseEnvelopeErrors]
type threatSignalFeedDeleteResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalFeedDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalFeedDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalFeedDeleteResponseEnvelopeMessages struct {
	Message  string                                               `json:"message" api:"required"`
	Code     float64                                              `json:"code"`
	Expected string                                               `json:"expected"`
	Path     []string                                             `json:"path"`
	Reason   ThreatSignalFeedDeleteResponseEnvelopeMessagesReason `json:"reason"`
	JSON     threatSignalFeedDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// threatSignalFeedDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ThreatSignalFeedDeleteResponseEnvelopeMessages]
type threatSignalFeedDeleteResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Expected    apijson.Field
	Path        apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalFeedDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalFeedDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalFeedDeleteResponseEnvelopeMessagesReason string

const (
	ThreatSignalFeedDeleteResponseEnvelopeMessagesReasonFreeCustomFeedLimit              ThreatSignalFeedDeleteResponseEnvelopeMessagesReason = "free_custom_feed_limit"
	ThreatSignalFeedDeleteResponseEnvelopeMessagesReasonFreeCustomSkillsDisabled         ThreatSignalFeedDeleteResponseEnvelopeMessagesReason = "free_custom_skills_disabled"
	ThreatSignalFeedDeleteResponseEnvelopeMessagesReasonFreeTierReconciliationInProgress ThreatSignalFeedDeleteResponseEnvelopeMessagesReason = "free_tier_reconciliation_in_progress"
	ThreatSignalFeedDeleteResponseEnvelopeMessagesReasonFreeTierReconciliationFailed     ThreatSignalFeedDeleteResponseEnvelopeMessagesReason = "free_tier_reconciliation_failed"
	ThreatSignalFeedDeleteResponseEnvelopeMessagesReasonRawContentExpired                ThreatSignalFeedDeleteResponseEnvelopeMessagesReason = "raw_content_expired"
	ThreatSignalFeedDeleteResponseEnvelopeMessagesReasonCuratedVisibilityChanged         ThreatSignalFeedDeleteResponseEnvelopeMessagesReason = "curated_visibility_changed"
)

func (r ThreatSignalFeedDeleteResponseEnvelopeMessagesReason) IsKnown() bool {
	switch r {
	case ThreatSignalFeedDeleteResponseEnvelopeMessagesReasonFreeCustomFeedLimit, ThreatSignalFeedDeleteResponseEnvelopeMessagesReasonFreeCustomSkillsDisabled, ThreatSignalFeedDeleteResponseEnvelopeMessagesReasonFreeTierReconciliationInProgress, ThreatSignalFeedDeleteResponseEnvelopeMessagesReasonFreeTierReconciliationFailed, ThreatSignalFeedDeleteResponseEnvelopeMessagesReasonRawContentExpired, ThreatSignalFeedDeleteResponseEnvelopeMessagesReasonCuratedVisibilityChanged:
		return true
	}
	return false
}

type ThreatSignalFeedDeleteResponseEnvelopeSuccess bool

const (
	ThreatSignalFeedDeleteResponseEnvelopeSuccessTrue ThreatSignalFeedDeleteResponseEnvelopeSuccess = true
)

func (r ThreatSignalFeedDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ThreatSignalFeedDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ThreatSignalFeedEditParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// One of the predefined Threat Signals feed categories; see GET
	// /:account_id/v2/threat-signals/categories.
	CategoryID    param.Field[ThreatSignalFeedEditParamsCategoryID] `json:"category_id"`
	DisplayName   param.Field[string]                               `json:"display_name"`
	Enabled       param.Field[bool]                                 `json:"enabled"`
	PollIntervalS param.Field[int64]                                `json:"poll_interval_s"`
	Title         param.Field[string]                               `json:"title"`
}

func (r ThreatSignalFeedEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// One of the predefined Threat Signals feed categories; see GET
// /:account_id/v2/threat-signals/categories.
type ThreatSignalFeedEditParamsCategoryID string

const (
	ThreatSignalFeedEditParamsCategoryIDB12a0fd6F7b9_5393_9ef3F888d506c550  ThreatSignalFeedEditParamsCategoryID = "b12a0fd6-f7b9-5393-9ef3-f888d506c550"
	ThreatSignalFeedEditParamsCategoryIDD5b70eaa626f5761B55b6d9590df49fb    ThreatSignalFeedEditParamsCategoryID = "d5b70eaa-626f-5761-b55b-6d9590df49fb"
	ThreatSignalFeedEditParamsCategoryID3b572d2b890d5286_9433F18c85079030   ThreatSignalFeedEditParamsCategoryID = "3b572d2b-890d-5286-9433-f18c85079030"
	ThreatSignalFeedEditParamsCategoryID17f90d3b37d3_5241_8ad4_7d6abbc2006c ThreatSignalFeedEditParamsCategoryID = "17f90d3b-37d3-5241-8ad4-7d6abbc2006c"
	ThreatSignalFeedEditParamsCategoryIDC68f28e9_7e8f5d4b853bF3076893a9ee   ThreatSignalFeedEditParamsCategoryID = "c68f28e9-7e8f-5d4b-853b-f3076893a9ee"
	ThreatSignalFeedEditParamsCategoryIDBb0e4a94_38ab5c14_80a7_28cee9f4b139 ThreatSignalFeedEditParamsCategoryID = "bb0e4a94-38ab-5c14-80a7-28cee9f4b139"
	ThreatSignalFeedEditParamsCategoryIDB1ef66d9A73c58dcB269_22d34dfd11f4   ThreatSignalFeedEditParamsCategoryID = "b1ef66d9-a73c-58dc-b269-22d34dfd11f4"
	ThreatSignalFeedEditParamsCategoryIDAb02a976_0a20_5c76A553_7f6325afacfe ThreatSignalFeedEditParamsCategoryID = "ab02a976-0a20-5c76-a553-7f6325afacfe"
)

func (r ThreatSignalFeedEditParamsCategoryID) IsKnown() bool {
	switch r {
	case ThreatSignalFeedEditParamsCategoryIDB12a0fd6F7b9_5393_9ef3F888d506c550, ThreatSignalFeedEditParamsCategoryIDD5b70eaa626f5761B55b6d9590df49fb, ThreatSignalFeedEditParamsCategoryID3b572d2b890d5286_9433F18c85079030, ThreatSignalFeedEditParamsCategoryID17f90d3b37d3_5241_8ad4_7d6abbc2006c, ThreatSignalFeedEditParamsCategoryIDC68f28e9_7e8f5d4b853bF3076893a9ee, ThreatSignalFeedEditParamsCategoryIDBb0e4a94_38ab5c14_80a7_28cee9f4b139, ThreatSignalFeedEditParamsCategoryIDB1ef66d9A73c58dcB269_22d34dfd11f4, ThreatSignalFeedEditParamsCategoryIDAb02a976_0a20_5c76A553_7f6325afacfe:
		return true
	}
	return false
}

type ThreatSignalFeedEditResponseEnvelope struct {
	Errors   []ThreatSignalFeedEditResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ThreatSignalFeedEditResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   ThreatSignalFeedEditResponse                   `json:"result" api:"required"`
	Success  ThreatSignalFeedEditResponseEnvelopeSuccess    `json:"success" api:"required"`
	JSON     threatSignalFeedEditResponseEnvelopeJSON       `json:"-"`
}

// threatSignalFeedEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [ThreatSignalFeedEditResponseEnvelope]
type threatSignalFeedEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalFeedEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalFeedEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalFeedEditResponseEnvelopeErrors struct {
	Message string                                         `json:"message" api:"required"`
	JSON    threatSignalFeedEditResponseEnvelopeErrorsJSON `json:"-"`
}

// threatSignalFeedEditResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ThreatSignalFeedEditResponseEnvelopeErrors]
type threatSignalFeedEditResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalFeedEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalFeedEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalFeedEditResponseEnvelopeMessages struct {
	Message  string                                             `json:"message" api:"required"`
	Code     float64                                            `json:"code"`
	Expected string                                             `json:"expected"`
	Path     []string                                           `json:"path"`
	Reason   ThreatSignalFeedEditResponseEnvelopeMessagesReason `json:"reason"`
	JSON     threatSignalFeedEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// threatSignalFeedEditResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ThreatSignalFeedEditResponseEnvelopeMessages]
type threatSignalFeedEditResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Expected    apijson.Field
	Path        apijson.Field
	Reason      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalFeedEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalFeedEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalFeedEditResponseEnvelopeMessagesReason string

const (
	ThreatSignalFeedEditResponseEnvelopeMessagesReasonFreeCustomFeedLimit              ThreatSignalFeedEditResponseEnvelopeMessagesReason = "free_custom_feed_limit"
	ThreatSignalFeedEditResponseEnvelopeMessagesReasonFreeCustomSkillsDisabled         ThreatSignalFeedEditResponseEnvelopeMessagesReason = "free_custom_skills_disabled"
	ThreatSignalFeedEditResponseEnvelopeMessagesReasonFreeTierReconciliationInProgress ThreatSignalFeedEditResponseEnvelopeMessagesReason = "free_tier_reconciliation_in_progress"
	ThreatSignalFeedEditResponseEnvelopeMessagesReasonFreeTierReconciliationFailed     ThreatSignalFeedEditResponseEnvelopeMessagesReason = "free_tier_reconciliation_failed"
	ThreatSignalFeedEditResponseEnvelopeMessagesReasonRawContentExpired                ThreatSignalFeedEditResponseEnvelopeMessagesReason = "raw_content_expired"
	ThreatSignalFeedEditResponseEnvelopeMessagesReasonCuratedVisibilityChanged         ThreatSignalFeedEditResponseEnvelopeMessagesReason = "curated_visibility_changed"
)

func (r ThreatSignalFeedEditResponseEnvelopeMessagesReason) IsKnown() bool {
	switch r {
	case ThreatSignalFeedEditResponseEnvelopeMessagesReasonFreeCustomFeedLimit, ThreatSignalFeedEditResponseEnvelopeMessagesReasonFreeCustomSkillsDisabled, ThreatSignalFeedEditResponseEnvelopeMessagesReasonFreeTierReconciliationInProgress, ThreatSignalFeedEditResponseEnvelopeMessagesReasonFreeTierReconciliationFailed, ThreatSignalFeedEditResponseEnvelopeMessagesReasonRawContentExpired, ThreatSignalFeedEditResponseEnvelopeMessagesReasonCuratedVisibilityChanged:
		return true
	}
	return false
}

type ThreatSignalFeedEditResponseEnvelopeSuccess bool

const (
	ThreatSignalFeedEditResponseEnvelopeSuccessTrue ThreatSignalFeedEditResponseEnvelopeSuccess = true
)

func (r ThreatSignalFeedEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ThreatSignalFeedEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ThreatSignalFeedPollParams struct {
	AccountID param.Field[string]                           `path:"account_id" api:"required"`
	FeedID    param.Field[ThreatSignalFeedPollParamsFeedID] `query:"feed_id"`
}

// URLQuery serializes [ThreatSignalFeedPollParams]'s query parameters as
// `url.Values`.
func (r ThreatSignalFeedPollParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ThreatSignalFeedPollParamsFeedID string

const (
	ThreatSignalFeedPollParamsFeedIDAll ThreatSignalFeedPollParamsFeedID = "all"
)

func (r ThreatSignalFeedPollParamsFeedID) IsKnown() bool {
	switch r {
	case ThreatSignalFeedPollParamsFeedIDAll:
		return true
	}
	return false
}

type ThreatSignalFeedPollResponseEnvelope struct {
	Errors  []ThreatSignalFeedPollResponseEnvelopeErrors `json:"errors" api:"required"`
	Result  ThreatSignalFeedPollResponse                 `json:"result" api:"required"`
	Success ThreatSignalFeedPollResponseEnvelopeSuccess  `json:"success" api:"required"`
	JSON    threatSignalFeedPollResponseEnvelopeJSON     `json:"-"`
}

// threatSignalFeedPollResponseEnvelopeJSON contains the JSON metadata for the
// struct [ThreatSignalFeedPollResponseEnvelope]
type threatSignalFeedPollResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalFeedPollResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalFeedPollResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalFeedPollResponseEnvelopeErrors struct {
	Message string                                         `json:"message" api:"required"`
	JSON    threatSignalFeedPollResponseEnvelopeErrorsJSON `json:"-"`
}

// threatSignalFeedPollResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ThreatSignalFeedPollResponseEnvelopeErrors]
type threatSignalFeedPollResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalFeedPollResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalFeedPollResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalFeedPollResponseEnvelopeSuccess bool

const (
	ThreatSignalFeedPollResponseEnvelopeSuccessTrue ThreatSignalFeedPollResponseEnvelopeSuccess = true
)

func (r ThreatSignalFeedPollResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ThreatSignalFeedPollResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
