// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_audit

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
	"github.com/cloudflare/cloudflare-go/v7/shared"
)

// RobotService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRobotService] method instead.
type RobotService struct {
	Options []option.RequestOption
}

// NewRobotService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRobotService(opts ...option.RequestOption) (r *RobotService) {
	r = &RobotService{}
	r.Options = opts
	return
}

// Fetches and parses robots.txt files for multiple domains within a zone in a
// single request. Each domain must belong to the specified zone. Results are keyed
// by hostname.
func (r *RobotService) BulkGet(ctx context.Context, params RobotBulkGetParams, opts ...option.RequestOption) (res *RobotBulkGetResponse, err error) {
	var env RobotBulkGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/ai-audit/robots/bulk", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Fetches and parses the robots.txt file for a zone or a specific subdomain within
// the zone. Returns parsed user-agent rules, content signals, and sitemaps.
func (r *RobotService) Get(ctx context.Context, params RobotGetParams, opts ...option.RequestOption) (res *RobotGetResponse, err error) {
	var env RobotGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/ai-audit/robots", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type RobotBulkGetResponse map[string]RobotBulkGetResponseItem

// Parsed robots.txt rules for a single domain.
type RobotBulkGetResponseItem struct {
	// Map of user-agent string to its parsed rules.
	UserAgents map[string]RobotBulkGetResponseItemUserAgent `json:"userAgents" api:"required"`
	// List of sitemap URLs found in robots.txt.
	Sitemaps []string `json:"sitemaps"`
	// HTTP status code from fetching the robots.txt file.
	Status int64                        `json:"status"`
	JSON   robotBulkGetResponseItemJSON `json:"-"`
}

// robotBulkGetResponseItemJSON contains the JSON metadata for the struct
// [RobotBulkGetResponseItem]
type robotBulkGetResponseItemJSON struct {
	UserAgents  apijson.Field
	Sitemaps    apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RobotBulkGetResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r robotBulkGetResponseItemJSON) RawJSON() string {
	return r.raw
}

// Parsed rules for a specific user-agent.
type RobotBulkGetResponseItemUserAgent struct {
	// List of allowed path patterns.
	Allow []string `json:"allow" api:"required"`
	// List of disallowed path patterns.
	Disallow []string `json:"disallow" api:"required"`
	// Content signal directives from robots.txt.
	ContentSignals RobotBulkGetResponseItemUserAgentsContentSignals `json:"contentSignals"`
	// Crawl delay in seconds.
	CrawlDelay float64                               `json:"crawlDelay"`
	JSON       robotBulkGetResponseItemUserAgentJSON `json:"-"`
}

// robotBulkGetResponseItemUserAgentJSON contains the JSON metadata for the struct
// [RobotBulkGetResponseItemUserAgent]
type robotBulkGetResponseItemUserAgentJSON struct {
	Allow          apijson.Field
	Disallow       apijson.Field
	ContentSignals apijson.Field
	CrawlDelay     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RobotBulkGetResponseItemUserAgent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r robotBulkGetResponseItemUserAgentJSON) RawJSON() string {
	return r.raw
}

// Content signal directives from robots.txt.
type RobotBulkGetResponseItemUserAgentsContentSignals struct {
	// Whether AI input usage is permitted.
	AIInput RobotBulkGetResponseItemUserAgentsContentSignalsAIInput `json:"ai-input"`
	// Whether AI training is permitted.
	AITrain RobotBulkGetResponseItemUserAgentsContentSignalsAITrain `json:"ai-train"`
	// Whether search indexing is permitted.
	Search RobotBulkGetResponseItemUserAgentsContentSignalsSearch `json:"search"`
	JSON   robotBulkGetResponseItemUserAgentsContentSignalsJSON   `json:"-"`
}

// robotBulkGetResponseItemUserAgentsContentSignalsJSON contains the JSON metadata
// for the struct [RobotBulkGetResponseItemUserAgentsContentSignals]
type robotBulkGetResponseItemUserAgentsContentSignalsJSON struct {
	AIInput     apijson.Field
	AITrain     apijson.Field
	Search      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RobotBulkGetResponseItemUserAgentsContentSignals) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r robotBulkGetResponseItemUserAgentsContentSignalsJSON) RawJSON() string {
	return r.raw
}

// Whether AI input usage is permitted.
type RobotBulkGetResponseItemUserAgentsContentSignalsAIInput string

const (
	RobotBulkGetResponseItemUserAgentsContentSignalsAIInputYes RobotBulkGetResponseItemUserAgentsContentSignalsAIInput = "yes"
	RobotBulkGetResponseItemUserAgentsContentSignalsAIInputNo  RobotBulkGetResponseItemUserAgentsContentSignalsAIInput = "no"
)

func (r RobotBulkGetResponseItemUserAgentsContentSignalsAIInput) IsKnown() bool {
	switch r {
	case RobotBulkGetResponseItemUserAgentsContentSignalsAIInputYes, RobotBulkGetResponseItemUserAgentsContentSignalsAIInputNo:
		return true
	}
	return false
}

// Whether AI training is permitted.
type RobotBulkGetResponseItemUserAgentsContentSignalsAITrain string

const (
	RobotBulkGetResponseItemUserAgentsContentSignalsAITrainYes RobotBulkGetResponseItemUserAgentsContentSignalsAITrain = "yes"
	RobotBulkGetResponseItemUserAgentsContentSignalsAITrainNo  RobotBulkGetResponseItemUserAgentsContentSignalsAITrain = "no"
)

func (r RobotBulkGetResponseItemUserAgentsContentSignalsAITrain) IsKnown() bool {
	switch r {
	case RobotBulkGetResponseItemUserAgentsContentSignalsAITrainYes, RobotBulkGetResponseItemUserAgentsContentSignalsAITrainNo:
		return true
	}
	return false
}

// Whether search indexing is permitted.
type RobotBulkGetResponseItemUserAgentsContentSignalsSearch string

const (
	RobotBulkGetResponseItemUserAgentsContentSignalsSearchYes RobotBulkGetResponseItemUserAgentsContentSignalsSearch = "yes"
	RobotBulkGetResponseItemUserAgentsContentSignalsSearchNo  RobotBulkGetResponseItemUserAgentsContentSignalsSearch = "no"
)

func (r RobotBulkGetResponseItemUserAgentsContentSignalsSearch) IsKnown() bool {
	switch r {
	case RobotBulkGetResponseItemUserAgentsContentSignalsSearchYes, RobotBulkGetResponseItemUserAgentsContentSignalsSearchNo:
		return true
	}
	return false
}

// Parsed robots.txt rules for a single domain.
type RobotGetResponse struct {
	// Map of user-agent string to its parsed rules.
	UserAgents map[string]RobotGetResponseUserAgent `json:"userAgents" api:"required"`
	// List of sitemap URLs found in robots.txt.
	Sitemaps []string `json:"sitemaps"`
	// HTTP status code from fetching the robots.txt file.
	Status int64                `json:"status"`
	JSON   robotGetResponseJSON `json:"-"`
}

// robotGetResponseJSON contains the JSON metadata for the struct
// [RobotGetResponse]
type robotGetResponseJSON struct {
	UserAgents  apijson.Field
	Sitemaps    apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RobotGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r robotGetResponseJSON) RawJSON() string {
	return r.raw
}

// Parsed rules for a specific user-agent.
type RobotGetResponseUserAgent struct {
	// List of allowed path patterns.
	Allow []string `json:"allow" api:"required"`
	// List of disallowed path patterns.
	Disallow []string `json:"disallow" api:"required"`
	// Content signal directives from robots.txt.
	ContentSignals RobotGetResponseUserAgentsContentSignals `json:"contentSignals"`
	// Crawl delay in seconds.
	CrawlDelay float64                       `json:"crawlDelay"`
	JSON       robotGetResponseUserAgentJSON `json:"-"`
}

// robotGetResponseUserAgentJSON contains the JSON metadata for the struct
// [RobotGetResponseUserAgent]
type robotGetResponseUserAgentJSON struct {
	Allow          apijson.Field
	Disallow       apijson.Field
	ContentSignals apijson.Field
	CrawlDelay     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RobotGetResponseUserAgent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r robotGetResponseUserAgentJSON) RawJSON() string {
	return r.raw
}

// Content signal directives from robots.txt.
type RobotGetResponseUserAgentsContentSignals struct {
	// Whether AI input usage is permitted.
	AIInput RobotGetResponseUserAgentsContentSignalsAIInput `json:"ai-input"`
	// Whether AI training is permitted.
	AITrain RobotGetResponseUserAgentsContentSignalsAITrain `json:"ai-train"`
	// Whether search indexing is permitted.
	Search RobotGetResponseUserAgentsContentSignalsSearch `json:"search"`
	JSON   robotGetResponseUserAgentsContentSignalsJSON   `json:"-"`
}

// robotGetResponseUserAgentsContentSignalsJSON contains the JSON metadata for the
// struct [RobotGetResponseUserAgentsContentSignals]
type robotGetResponseUserAgentsContentSignalsJSON struct {
	AIInput     apijson.Field
	AITrain     apijson.Field
	Search      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RobotGetResponseUserAgentsContentSignals) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r robotGetResponseUserAgentsContentSignalsJSON) RawJSON() string {
	return r.raw
}

// Whether AI input usage is permitted.
type RobotGetResponseUserAgentsContentSignalsAIInput string

const (
	RobotGetResponseUserAgentsContentSignalsAIInputYes RobotGetResponseUserAgentsContentSignalsAIInput = "yes"
	RobotGetResponseUserAgentsContentSignalsAIInputNo  RobotGetResponseUserAgentsContentSignalsAIInput = "no"
)

func (r RobotGetResponseUserAgentsContentSignalsAIInput) IsKnown() bool {
	switch r {
	case RobotGetResponseUserAgentsContentSignalsAIInputYes, RobotGetResponseUserAgentsContentSignalsAIInputNo:
		return true
	}
	return false
}

// Whether AI training is permitted.
type RobotGetResponseUserAgentsContentSignalsAITrain string

const (
	RobotGetResponseUserAgentsContentSignalsAITrainYes RobotGetResponseUserAgentsContentSignalsAITrain = "yes"
	RobotGetResponseUserAgentsContentSignalsAITrainNo  RobotGetResponseUserAgentsContentSignalsAITrain = "no"
)

func (r RobotGetResponseUserAgentsContentSignalsAITrain) IsKnown() bool {
	switch r {
	case RobotGetResponseUserAgentsContentSignalsAITrainYes, RobotGetResponseUserAgentsContentSignalsAITrainNo:
		return true
	}
	return false
}

// Whether search indexing is permitted.
type RobotGetResponseUserAgentsContentSignalsSearch string

const (
	RobotGetResponseUserAgentsContentSignalsSearchYes RobotGetResponseUserAgentsContentSignalsSearch = "yes"
	RobotGetResponseUserAgentsContentSignalsSearchNo  RobotGetResponseUserAgentsContentSignalsSearch = "no"
)

func (r RobotGetResponseUserAgentsContentSignalsSearch) IsKnown() bool {
	switch r {
	case RobotGetResponseUserAgentsContentSignalsSearchYes, RobotGetResponseUserAgentsContentSignalsSearchNo:
		return true
	}
	return false
}

type RobotBulkGetParams struct {
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
	// Array of domain hostnames to fetch robots.txt for. Each domain must end with the
	// zone name. Maximum 25 domains per request.
	Body []string `json:"body" api:"required"`
}

func (r RobotBulkGetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type RobotBulkGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	Success  bool                  `json:"success" api:"required"`
	// Map of hostname to parsed robots.txt rules.
	Result RobotBulkGetResponse             `json:"result"`
	JSON   robotBulkGetResponseEnvelopeJSON `json:"-"`
}

// robotBulkGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RobotBulkGetResponseEnvelope]
type robotBulkGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RobotBulkGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r robotBulkGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RobotGetParams struct {
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
	// Optional subdomain to fetch robots.txt for. If omitted, fetches robots.txt for
	// the zone apex domain.
	Subdomain param.Field[string] `query:"subdomain"`
}

// URLQuery serializes [RobotGetParams]'s query parameters as `url.Values`.
func (r RobotGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type RobotGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	Success  bool                  `json:"success" api:"required"`
	// Parsed robots.txt rules for a single domain.
	Result RobotGetResponse             `json:"result"`
	JSON   robotGetResponseEnvelopeJSON `json:"-"`
}

// robotGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RobotGetResponseEnvelope]
type robotGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RobotGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r robotGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
