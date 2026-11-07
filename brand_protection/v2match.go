// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brand_protection

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

// V2MatchService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2MatchService] method instead.
type V2MatchService struct {
	Options []option.RequestOption
}

// NewV2MatchService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV2MatchService(opts ...option.RequestOption) (r *V2MatchService) {
	r = &V2MatchService{}
	r.Options = opts
	return
}

// Lists domains that matched one or more saved domain queries, most recently
// matched first by default. When you pass several comma-separated `query_id`
// values, matches are deduplicated by domain and each match includes a
// `match_details` array with per-query dismissed state. Set
// `include_domain_id=true` to get the `domain_id` needed to dismiss or undismiss a
// match.
func (r *V2MatchService) Get(ctx context.Context, params V2MatchGetParams, opts ...option.RequestOption) (res *V2MatchGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/v2/brand-protection/domain/matches", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

type V2MatchGetResponse struct {
	Matches []V2MatchGetResponseMatch `json:"matches" api:"required"`
	// Number of matching domains available for pagination. For requests with several
	// `query_id` values, matches are deduplicated from at most 5,000 query-domain
	// rows, so `total` can be lower than the actual number of matches.
	Total int64                  `json:"total" api:"required"`
	JSON  v2MatchGetResponseJSON `json:"-"`
}

// v2MatchGetResponseJSON contains the JSON metadata for the struct
// [V2MatchGetResponse]
type v2MatchGetResponseJSON struct {
	Matches     apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2MatchGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2MatchGetResponseJSON) RawJSON() string {
	return r.raw
}

type V2MatchGetResponseMatch struct {
	// The matched domain name.
	Domain string `json:"domain" api:"required"`
	// When the domain was matched. For multi-query requests, the most recent match
	// time across the requested queries.
	FirstSeen   string                               `json:"first_seen" api:"required"`
	PublicScans V2MatchGetResponseMatchesPublicScans `json:"public_scans" api:"required,nullable"`
	// Registrar of the matched domain, if known.
	Registrar        string `json:"registrar" api:"required,nullable"`
	ScanStatus       string `json:"scan_status" api:"required"`
	ScanSubmissionID int64  `json:"scan_submission_id" api:"required,nullable"`
	Source           string `json:"source" api:"required,nullable"`
	// Whether the match is dismissed. Only present for single-query requests. For
	// multi-query requests, use the dismissed field in each match_details entry.
	Dismissed bool `json:"dismissed"`
	// ID of the matched domain. Present only when `include_domain_id=true`. Use it
	// with `query_id` to dismiss or undismiss the match.
	DomainID int64 `json:"domain_id"`
	// Per-match detail objects with query metadata and individual dismissed state.
	// Only present when multiple query_ids are requested.
	MatchDetails []V2MatchGetResponseMatchesMatchDetail `json:"match_details"`
	JSON         v2MatchGetResponseMatchJSON            `json:"-"`
}

// v2MatchGetResponseMatchJSON contains the JSON metadata for the struct
// [V2MatchGetResponseMatch]
type v2MatchGetResponseMatchJSON struct {
	Domain           apijson.Field
	FirstSeen        apijson.Field
	PublicScans      apijson.Field
	Registrar        apijson.Field
	ScanStatus       apijson.Field
	ScanSubmissionID apijson.Field
	Source           apijson.Field
	Dismissed        apijson.Field
	DomainID         apijson.Field
	MatchDetails     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *V2MatchGetResponseMatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2MatchGetResponseMatchJSON) RawJSON() string {
	return r.raw
}

type V2MatchGetResponseMatchesPublicScans struct {
	SubmissionID string                                   `json:"submission_id" api:"required"`
	JSON         v2MatchGetResponseMatchesPublicScansJSON `json:"-"`
}

// v2MatchGetResponseMatchesPublicScansJSON contains the JSON metadata for the
// struct [V2MatchGetResponseMatchesPublicScans]
type v2MatchGetResponseMatchesPublicScansJSON struct {
	SubmissionID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *V2MatchGetResponseMatchesPublicScans) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2MatchGetResponseMatchesPublicScansJSON) RawJSON() string {
	return r.raw
}

type V2MatchGetResponseMatchesMatchDetail struct {
	// Individual dismissed state for this specific match.
	Dismissed bool `json:"dismissed" api:"required"`
	// ID of the match. Accepted by the legacy `match_ids` format.
	MatchID int64 `json:"match_id" api:"required"`
	// ID of the domain query that produced this match.
	QueryID int64 `json:"query_id" api:"required"`
	// Tag associated with the query, if one exists.
	QueryTag string                                   `json:"query_tag" api:"required,nullable"`
	JSON     v2MatchGetResponseMatchesMatchDetailJSON `json:"-"`
}

// v2MatchGetResponseMatchesMatchDetailJSON contains the JSON metadata for the
// struct [V2MatchGetResponseMatchesMatchDetail]
type v2MatchGetResponseMatchesMatchDetailJSON struct {
	Dismissed   apijson.Field
	MatchID     apijson.Field
	QueryID     apijson.Field
	QueryTag    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2MatchGetResponseMatchesMatchDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2MatchGetResponseMatchesMatchDetailJSON) RawJSON() string {
	return r.raw
}

type V2MatchGetParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// ID of a domain query, or a comma-separated list of domain query IDs. When you
	// pass several IDs, matches are deduplicated by domain and each match includes a
	// `match_details` array with per-query metadata and dismissed state.
	QueryID param.Field[[]string] `query:"query_id" api:"required"`
	// Return only matches whose domain contains this text (case-insensitive).
	DomainSearch param.Field[string] `query:"domain_search"`
	// Set to `false` to exclude dismissed matches. Dismissed matches are included when
	// this is omitted or `true`.
	IncludeDismissed param.Field[string] `query:"include_dismissed"`
	// Set to `true` to include `domain_id` on each match. Required to dismiss or
	// undismiss a match. Defaults to `false`.
	IncludeDomainID param.Field[string] `query:"include_domain_id"`
	// Maximum number of matches to return. Defaults to 50.
	Limit param.Field[string] `query:"limit"`
	// Number of matches to skip, for pagination. Defaults to 0.
	Offset param.Field[string] `query:"offset"`
	// Sort direction: `asc` or `desc`. Defaults to `desc`. For single-query requests,
	// `asc` applies only when `orderBy` is set.
	Order param.Field[V2MatchGetParamsOrder] `query:"order"`
	// Field to sort by: `domain`, `first_seen`, or `registrar`. Defaults to
	// `first_seen`.
	OrderBy param.Field[V2MatchGetParamsOrderBy] `query:"orderBy"`
}

// URLQuery serializes [V2MatchGetParams]'s query parameters as `url.Values`.
func (r V2MatchGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Sort direction: `asc` or `desc`. Defaults to `desc`. For single-query requests,
// `asc` applies only when `orderBy` is set.
type V2MatchGetParamsOrder string

const (
	V2MatchGetParamsOrderAsc  V2MatchGetParamsOrder = "asc"
	V2MatchGetParamsOrderDesc V2MatchGetParamsOrder = "desc"
)

func (r V2MatchGetParamsOrder) IsKnown() bool {
	switch r {
	case V2MatchGetParamsOrderAsc, V2MatchGetParamsOrderDesc:
		return true
	}
	return false
}

// Field to sort by: `domain`, `first_seen`, or `registrar`. Defaults to
// `first_seen`.
type V2MatchGetParamsOrderBy string

const (
	V2MatchGetParamsOrderByDomain    V2MatchGetParamsOrderBy = "domain"
	V2MatchGetParamsOrderByFirstSeen V2MatchGetParamsOrderBy = "first_seen"
	V2MatchGetParamsOrderByRegistrar V2MatchGetParamsOrderBy = "registrar"
)

func (r V2MatchGetParamsOrderBy) IsKnown() bool {
	switch r {
	case V2MatchGetParamsOrderByDomain, V2MatchGetParamsOrderByFirstSeen, V2MatchGetParamsOrderByRegistrar:
		return true
	}
	return false
}
