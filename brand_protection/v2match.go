// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brand_protection

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
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

// Get paginated list of domain matches for one or more brand protection queries.
// When multiple query_ids are provided (comma-separated), matches are deduplicated
// across queries and each match includes a match_details array with per-match
// query metadata and individual dismissed state.
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
	Total   int64                     `json:"total" api:"required"`
	JSON    v2MatchGetResponseJSON    `json:"-"`
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
	Domain           string                               `json:"domain" api:"required"`
	FirstSeen        string                               `json:"first_seen" api:"required"`
	PublicScans      V2MatchGetResponseMatchesPublicScans `json:"public_scans" api:"required,nullable"`
	Registrar        string                               `json:"registrar" api:"required,nullable"`
	ScanStatus       string                               `json:"scan_status" api:"required"`
	ScanSubmissionID int64                                `json:"scan_submission_id" api:"required,nullable"`
	Source           string                               `json:"source" api:"required,nullable"`
	// Whether the match is dismissed. Only present for single-query requests. For
	// multi-query requests, use the dismissed field in each match_details entry.
	Dismissed bool `json:"dismissed"`
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
	Dismissed bool  `json:"dismissed" api:"required"`
	MatchID   int64 `json:"match_id" api:"required"`
	QueryID   int64 `json:"query_id" api:"required"`
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
	// Query ID or comma-separated list of Query IDs. When multiple IDs are provided,
	// matches are deduplicated across queries and each match includes a match_details
	// array with per-match query metadata and dismissed state.
	QueryID param.Field[[]string] `query:"query_id" api:"required"`
	// Filter matches by domain name (substring match)
	DomainSearch     param.Field[string] `query:"domain_search"`
	IncludeDismissed param.Field[string] `query:"include_dismissed"`
	IncludeDomainID  param.Field[string] `query:"include_domain_id"`
	Limit            param.Field[string] `query:"limit"`
	Offset           param.Field[string] `query:"offset"`
	// Sort order. Options: 'asc' (ascending) or 'desc' (descending)
	Order param.Field[V2MatchGetParamsOrder] `query:"order"`
	// Column to sort by. Options: 'domain', 'first_seen', or 'registrar'
	OrderBy param.Field[V2MatchGetParamsOrderBy] `query:"orderBy"`
}

// URLQuery serializes [V2MatchGetParams]'s query parameters as `url.Values`.
func (r V2MatchGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Sort order. Options: 'asc' (ascending) or 'desc' (descending)
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

// Column to sort by. Options: 'domain', 'first_seen', or 'registrar'
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
