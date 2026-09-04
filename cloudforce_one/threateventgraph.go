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
)

// ThreatEventGraphService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreatEventGraphService] method instead.
type ThreatEventGraphService struct {
	Options []option.RequestOption
}

// NewThreatEventGraphService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewThreatEventGraphService(opts ...option.RequestOption) (r *ThreatEventGraphService) {
	r = &ThreatEventGraphService{}
	r.Options = opts
	return
}

// Expands the single-level relationship neighborhood of one or more seed nodes
// (event, indicator, or tag) from R2 Data Catalog. Seeds use compact id format
// (type:uuid), e.g. "event:550e8400-...". Multi-seed requests merge and
// deduplicate results server-side. Hydrates neighbor entities with summary data
// from Durable Objects. Supports filtering by relationship type and dataset scope.
func (r *ThreatEventGraphService) List(ctx context.Context, params ThreatEventGraphListParams, opts ...option.RequestOption) (res *ThreatEventGraphListResponse, err error) {
	var env ThreatEventGraphListResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/graph", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type ThreatEventGraphListResponse struct {
	Edges []ThreatEventGraphListResponseEdge `json:"edges" api:"required"`
	// Focal node object (legacy single-seed). Null when unavailable.
	Node  map[string]interface{}           `json:"node" api:"required,nullable"`
	Nodes []map[string]interface{}         `json:"nodes" api:"required"`
	JSON  threatEventGraphListResponseJSON `json:"-"`
}

// threatEventGraphListResponseJSON contains the JSON metadata for the struct
// [ThreatEventGraphListResponse]
type threatEventGraphListResponseJSON struct {
	Edges       apijson.Field
	Node        apijson.Field
	Nodes       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventGraphListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventGraphListResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventGraphListResponseEdge struct {
	// Deterministic composite edge id (source→target:relationshipType)
	ID               string `json:"id" api:"required"`
	RelationshipType string `json:"relationshipType" api:"required"`
	// Compact id of the source node (type:uuid)
	Source     string `json:"source" api:"required"`
	SourceID   string `json:"sourceId" api:"required"`
	SourceType string `json:"sourceType" api:"required"`
	// Compact id of the target node (type:uuid)
	Target     string                               `json:"target" api:"required"`
	TargetID   string                               `json:"targetId" api:"required"`
	TargetType string                               `json:"targetType" api:"required"`
	JSON       threatEventGraphListResponseEdgeJSON `json:"-"`
}

// threatEventGraphListResponseEdgeJSON contains the JSON metadata for the struct
// [ThreatEventGraphListResponseEdge]
type threatEventGraphListResponseEdgeJSON struct {
	ID               apijson.Field
	RelationshipType apijson.Field
	Source           apijson.Field
	SourceID         apijson.Field
	SourceType       apijson.Field
	Target           apijson.Field
	TargetID         apijson.Field
	TargetType       apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ThreatEventGraphListResponseEdge) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventGraphListResponseEdgeJSON) RawJSON() string {
	return r.raw
}

type ThreatEventGraphListParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Opaque pagination token. Only valid when seeds has exactly 1 entry; 400
	// otherwise.
	Cursor param.Field[string] `query:"cursor"`
	// Comma-separated dataset UUIDs to restrict neighbor scope, or one standalone
	// scope value: 'all'/'\*', 'analytics' for isAnalytics=true datasets, or
	// 'operational' for isAnalytics=false datasets. Intersected with access grants.
	DatasetIDs param.Field[[]string] `query:"datasetIds"`
	// Edge direction relative to each seed: out (seed→neighbors), in (neighbors→seed),
	// both (default).
	Direction param.Field[string] `query:"direction"`
	// Comma-separated list of response sections to expand (hydrate). Allowed: `nodes`.
	// Omitting `expand` returns identifier-only nodes.
	Expand param.Field[[]string] `query:"expand"`
	// Hydration strategy for neighbor nodes when expand=nodes is set. r2_join
	// (default): use R2 JOIN query + DO fallback. do_only: use plain R2 query +
	// hydrate all neighbors via Durable Objects.
	Hydration param.Field[string] `query:"hydration"`
	// Max neighbors per seed (default: 100, max: 1000). Values above 1000 return 400.
	Limit param.Field[float64] `query:"limit"`
	// Total accumulated node cap across all seeds (default: 500, max: 1000). Values
	// above 1000 return 400.
	MaxNodes param.Field[float64] `query:"max_nodes"`
	// Comma-separated relationship types to filter by. Allowed: tagged_with,
	// appears_in, related_to, caused_by, attributed_to.
	RelationshipTypes param.Field[[]string] `query:"relationshipTypes"`
	// Comma-separated compact seed ids (type:uuid). Example:
	// seeds=event:550e8400-...,indicator:661fa920-... Provide 1–50 entries; omitting
	// seeds returns 400.
	Seeds param.Field[[]string] `query:"seeds"`
}

// URLQuery serializes [ThreatEventGraphListParams]'s query parameters as
// `url.Values`.
func (r ThreatEventGraphListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ThreatEventGraphListResponseEnvelope struct {
	Errors     []map[string]interface{}                       `json:"errors" api:"required"`
	Messages   []map[string]interface{}                       `json:"messages" api:"required"`
	Result     ThreatEventGraphListResponse                   `json:"result" api:"required"`
	Success    bool                                           `json:"success" api:"required"`
	ResultInfo ThreatEventGraphListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       threatEventGraphListResponseEnvelopeJSON       `json:"-"`
}

// threatEventGraphListResponseEnvelopeJSON contains the JSON metadata for the
// struct [ThreatEventGraphListResponseEnvelope]
type threatEventGraphListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventGraphListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventGraphListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ThreatEventGraphListResponseEnvelopeResultInfo struct {
	// Number of nodes in result.nodes (seeds + neighbors)
	Count float64 `json:"count" api:"required"`
	// Number of edges in result.edges
	EdgeCount float64 `json:"edge_count" api:"required"`
	// Total query time in milliseconds
	QueryTimeMs float64 `json:"query_time_ms" api:"required"`
	// Total count of nodes (same as count for this endpoint)
	TotalCount float64 `json:"total_count" api:"required"`
	// Opaque pagination cursor for the next page; null when exhausted or for
	// multi-seed requests (single-seed only)
	Cursor string `json:"cursor" api:"nullable"`
	// Traversal depth reached (always 1 for single-level)
	DepthReached float64 `json:"depth_reached"`
	// True when a cursor is available for the next page (single-seed only)
	HasMore bool `json:"has_more"`
	// Composite ids of the seed node(s) (type:uuid). Always an array, even for one
	// seed.
	Seeds []string `json:"seeds"`
	// True when results were capped (per-seed limit or max_nodes)
	Truncated bool                                               `json:"truncated"`
	JSON      threatEventGraphListResponseEnvelopeResultInfoJSON `json:"-"`
}

// threatEventGraphListResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [ThreatEventGraphListResponseEnvelopeResultInfo]
type threatEventGraphListResponseEnvelopeResultInfoJSON struct {
	Count        apijson.Field
	EdgeCount    apijson.Field
	QueryTimeMs  apijson.Field
	TotalCount   apijson.Field
	Cursor       apijson.Field
	DepthReached apijson.Field
	HasMore      apijson.Field
	Seeds        apijson.Field
	Truncated    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ThreatEventGraphListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventGraphListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
